package analyzer

import (
	"fmt"
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/parser/ast"
	"github/suixinpr/manadb/internal/parser/types"
	"github/suixinpr/manadb/internal/relation/catalog"
)

func (ana *Analyzer) transformExprNodeList(nodes []ast.ExprNode) ([]expression.Expression, error) {
	exprs := make([]expression.Expression, len(nodes))
	for i, node := range nodes {
		expr, err := ana.transformExprNode(node)
		if err != nil {
			return nil, err
		}
		exprs[i] = expr
	}
	return exprs, nil
}

func (ana *Analyzer) transformExprNode(node ast.ExprNode) (expression.Expression, error) {
	switch n := node.(type) {
	case nil:
		return nil, nil
	case *ast.ColumnRefExpr:
		return ana.transformColumnRefExpr(n)
	case *ast.ConstExpr:
		return ana.transformConstExpr(n)
	case *ast.OperationExpr:
		return ana.transformOperationExpr(n)
	default:
		return nil, errlog.New(fmt.Sprintf("transformExprNode not support ExprNode type %T", n))
	}
}

/* 将所有列引用和对应的表建立联系 */
func (ana *Analyzer) transformColumnRefExpr(node *ast.ColumnRefExpr) (*expression.Variable, error) {
	var result *expression.Variable
	switch {
	case node.TableName == "":
		for _, ns := range ana.namespace {
			tbl := ns.table
			for j, column := range tbl.Columns {
				if node.ColumnName != column.Name {
					continue
				}
				if result != nil {
					return nil, fmt.Errorf("column reference %s is ambiguous", node.ColumnName)
				}
				result = &expression.Variable{
					VarNo:    ns.index,
					VarColNo: j,
					VarType:  column.TypeID,
				}
			}
		}
	default:
		for _, ns := range ana.namespace {
			tbl := ns.table
			if node.TableName != tbl.TableName {
				continue
			}
			for j, column := range tbl.Columns {
				if node.ColumnName != column.Name {
					continue
				}
				result = &expression.Variable{
					VarNo:    ns.index,
					VarColNo: j,
					VarType:  column.TypeID,
				}
				break
			}
			break
		}
	}
	if result == nil {
		return nil, errlog.New("transformColumnRefExpr failed")
	}
	return result, nil
}

func (ana *Analyzer) transformConstExpr(node *ast.ConstExpr) (*expression.Const, error) {
	switch node.Type {
	case types.ConstNull:
		return &expression.Const{ConstType: metadata.InvalidID, ConstValue: nil}, nil
	case types.ConstBoolean:
		return &expression.Const{ConstType: metadata.BooleanID, ConstValue: datum.ValueGetDatum(node.B)}, nil
	case types.ConstInt:
		return &expression.Const{ConstType: metadata.Int64ID, ConstValue: datum.ValueGetDatum(node.I)}, nil
	case types.ConstUint:
		return &expression.Const{ConstType: metadata.Uint64ID, ConstValue: datum.ValueGetDatum(node.U)}, nil
	case types.ConstFloat:
		return &expression.Const{ConstType: metadata.Float64ID, ConstValue: datum.ValueGetDatum(node.F)}, nil
	case types.ConstString:
		return &expression.Const{ConstType: metadata.TextID, ConstValue: datum.StringGetDatum(node.S)}, nil
	default:
		return nil, errlog.New(fmt.Sprintf("transformConstExpr not support ConstType type %d", types.ConstNull))
	}
}

func (ana *Analyzer) transformOperationExpr(node *ast.OperationExpr) (*expression.Operation, error) {
	left, err := ana.transformExprNode(node.L)
	if err != nil {
		return nil, err
	}
	right, err := ana.transformExprNode(node.R)
	if err != nil {
		return nil, err
	}

	/* 获取操作符 */
	meta, err := catalog.SearchCatalog(metadata.ManaOperatorsID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaOperatorsOprName, metadata.TextID, datum.StringGetDatum(node.OprName)),
		catalog.NewScanKey(metadata.ManaOperatorsOprLeft, metadata.OIDID, datum.ValueGetDatum(left.TypeID())),
		catalog.NewScanKey(metadata.ManaOperatorsOprRight, metadata.OIDID, datum.ValueGetDatum(right.TypeID())),
	})
	if err != nil {
		return nil, err
	}

	/* 隐式类型转换 */
	if meta == nil {
		left, right, err = ana.castImplicitType(left, right)
		if err != nil {
			return nil, err
		}
	}

	/* 获取新的操作符 */
	meta, err = catalog.SearchCatalog(metadata.ManaOperatorsID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaOperatorsOprName, metadata.TextID, datum.StringGetDatum(node.OprName)),
		catalog.NewScanKey(metadata.ManaOperatorsOprLeft, metadata.OIDID, datum.ValueGetDatum(left.TypeID())),
		catalog.NewScanKey(metadata.ManaOperatorsOprRight, metadata.OIDID, datum.ValueGetDatum(right.TypeID())),
	})
	if err != nil {
		return nil, err
	}

	return &expression.Operation{
		OprID:     datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaOperatorsOprID]),
		OprLeft:   left,
		OprRight:  right,
		OprResult: datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaOperatorsOprResult]),
		OprFuncID: datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaOperatorsOprFuncID]),
	}, nil
}

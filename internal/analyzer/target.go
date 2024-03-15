package analyzer

import (
	"fmt"
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/analyzer/logic"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/parser/ast"
)

func (ana *Analyzer) transformFields(fields []ast.Node) ([]*logic.TargetInfo, error) {
	var targetList []*logic.TargetInfo
	for _, field := range fields {
		switch x := field.(type) {
		case *ast.FieldStar:
			for _, ns := range ana.namespace {
				tbl := ns.table
				if x.TableName != "" && x.TableName != tbl.TableName {
					continue
				}
				for _, col := range tbl.Columns {
					expr, err := ana.transformExprNode(&ast.ColumnRefExpr{
						TableName:  tbl.TableName,
						ColumnName: col.Name,
					})
					if err != nil {
						return nil, err
					}
					targetList = append(targetList, &logic.TargetInfo{
						Name: col.Name,
						Expr: expr,
					})
				}
			}
		case *ast.FieldColumn:
			expr, err := ana.transformExprNode(x.Expr)
			if err != nil {
				return nil, err
			}
			targetList = append(targetList, &logic.TargetInfo{
				Name: FieldColumnName(x),
				Expr: expr,
			})
		default:
			return nil, errlog.New(fmt.Sprintf("transformFields not support Node type %T", x))
		}
	}
	return targetList, nil
}

func FieldColumnName(n *ast.FieldColumn) string {
	if n.Name != "" {
		return n.Name
	}
	switch x := n.Expr.(type) {
	case *ast.ColumnRefExpr:
		return x.ColumnName
	default:
		return "?column?"
	}
}

func (ana *Analyzer) transformInsertValues(obj *logic.TableObject, valuesList [][]ast.ExprNode) ([][]expression.Expression, error) {
	tbl := ana.tables[obj.Index]
	exprsList := make([][]expression.Expression, len(valuesList))
	for i, values := range valuesList {
		if len(values) != len(tbl.Columns) {
			return nil, errlog.New("values column num and column num do not match")
		}
		exprs, err := ana.transformExprNodeList(values)
		if err != nil {
			return nil, err
		}
		exprsList[i] = exprs
	}

	/* 查找是否有变量 */
	var err error
	var walk expression.Walker
	walk = func(expr expression.Expression) {
		if err != nil || expr == nil {
			return
		}
		switch expr := expr.(type) {
		case *expression.Variable:
			err = errlog.New("columns are prohibited in insert values")
		default:
			expr.Walker(walk)
		}
	}
	for _, exprs := range exprsList {
		for _, expr := range exprs {
			walk(expr)
			if err != nil {
				return nil, err
			}
		}
	}

	/* 统一表和 values 的类型 */
	for i, exprs := range exprsList {
		for j, expr := range exprs {
			exprsList[i][j], err = ana.castToTargetType(expr, tbl.Columns[j].TypeID)
			if err != nil {
				return nil, err
			}
		}
	}

	return exprsList, err
}

func (ana *Analyzer) transformSetList(obj *logic.TableObject, setList []*ast.Assignment) ([]expression.Expression, error) {
	tbl := ana.tables[obj.Index]
	exprs := make([]expression.Expression, len(tbl.Columns))
	for _, assign := range setList {
		col, err := ana.transformColumnRefExpr(assign.ColumnRef)
		if err != nil {
			return nil, err
		}
		expr, err := ana.transformExprNode(assign.Expr)
		if err != nil {
			return nil, err
		}
		if exprs[col.VarColNo] != nil {
			return nil, errlog.New("update duplicate column")
		}
		exprs[col.VarColNo] = expr
	}

	/* 补充缺省值 */
	for i, expr := range exprs {
		if expr != nil {
			continue
		}
		exprs[i] = &expression.Variable{
			VarNo:    obj.Index,
			VarColNo: i,
			VarType:  tbl.Columns[i].TypeID,
		}
	}

	/* 统一表和 values 的类型 */
	var err error
	for i, expr := range exprs {
		exprs[i], err = ana.castToTargetType(expr, tbl.Columns[i].TypeID)
		if err != nil {
			return nil, err
		}
	}

	return exprs, err
}

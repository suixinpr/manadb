package analyzer

import (
	"fmt"
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/analyzer/logic"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/parser/ast"
	"github/suixinpr/manadb/internal/relation/catalog"
)

func Analyze(stmt ast.StmtNode) (logic.LogicPlan, error) {
	ana := new(Analyzer)
	return ana.analyze(stmt)
}

type Analyzer struct {
	namespace []*parseNamespace
	tables    []*logic.TableInfo
}

func (ana *Analyzer) analyze(stmt ast.StmtNode) (logic.LogicPlan, error) {
	switch x := stmt.(type) {
	case *ast.CreateTableStmt:
		return ana.transformCreateTableStmt(x)
	case *ast.DeleteStmt:
		return ana.transformDeleteStmt(x)
	case *ast.DropTableStmt:
		return ana.transformDropTableStmt(x)
	case *ast.ExplainStmt:
		return ana.transformExplainStmt(x)
	case *ast.InsertStmt:
		return ana.transformInsertStmt(x)
	case *ast.SelectStmt:
		return ana.transformSelectStmt(x)
	case *ast.UpdateStmt:
		return ana.transformUpdateStmt(x)
	default:
		return nil, errlog.New(fmt.Sprintf("Analyze not support StmtNode type %T", x))
	}
}

func (ana *Analyzer) transformCreateTableStmt(stmt *ast.CreateTableStmt) (*logic.CreateTablePlan, error) {
	plan := &logic.CreateTablePlan{TableName: stmt.TableName}
	for _, ele := range stmt.TableElementList {
		typ := ele.ColumnType
		meta, err := catalog.SearchCatalogOne(metadata.ManaTypesID, []*common.ScanKey{
			catalog.NewScanKey(metadata.ManaTypesTypeName, metadata.TextID, datum.StringGetDatum(typ.Name)),
		})
		if err != nil {
			return nil, err
		}
		plan.Columns = append(plan.Columns, &logic.ColumnInfo{
			Name:    ele.ColumnName,
			Mod:     int16(typ.Len),
			TypeID:  datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaTypesTypeID]),
			TypeLen: datum.DatumGetValue[int16](meta.Values[metadata.ManaTypesTypeLen]),
		})
	}
	return plan, nil
}

func (ana *Analyzer) transformDeleteStmt(stmt *ast.DeleteStmt) (*logic.DeletePlan, error) {
	/* 生成查询树 */
	query := new(logic.QueryPlan)

	from, err := ana.transformSpecTable(stmt.From)
	if err != nil {
		return nil, err
	}

	query.Qual, err = ana.transformWhereClause(stmt.WhereClause)
	if err != nil {
		return nil, err
	}

	query.Tables = ana.tables
	query.From = from

	/* 生成 delete plan */
	plan := new(logic.DeletePlan)
	plan.Table = ana.tables[from.Index]
	plan.Query = query
	return plan, nil
}

func (ana *Analyzer) transformDropTableStmt(stmt *ast.DropTableStmt) (*logic.DropTablePlan, error) {
	plan := new(logic.DropTablePlan)
	plan.Tables = make([]metadata.OID, len(stmt.Tables))

	for i, table := range stmt.Tables {
		name := table.TableName
		/* 检查表是否不存在 */
		meta, err := catalog.SearchCatalog(metadata.ManaTablesID, []*common.ScanKey{
			catalog.NewScanKey(metadata.ManaTablesTblName, metadata.TextID, datum.StringGetDatum(name)),
		})
		if err != nil {
			return nil, err
		}
		if meta == nil {
			return nil, errlog.New(fmt.Sprintf("table %s does not exist", name))
		}
		/* 获取表oid */
		tblID := datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaTablesTblID])
		if metadata.IsCatalog(tblID) {
			return nil, errlog.New(fmt.Sprintf("%s is a system catalog", name))
		}
		plan.Tables[i] = tblID
	}

	return plan, nil
}

func (ana *Analyzer) transformExplainStmt(stmt *ast.ExplainStmt) (*logic.ExplainPlan, error) {
	plan, err := ana.analyze(stmt.Stmt)
	if err != nil {
		return nil, err
	}
	return &logic.ExplainPlan{Plan: plan}, nil
}

func (ana *Analyzer) transformInsertStmt(stmt *ast.InsertStmt) (*logic.InsertPlan, error) {
	plan := new(logic.InsertPlan)

	tbl, err := ana.transformSpecTable(stmt.Table)
	if err != nil {
		return nil, err
	}
	plan.Table = ana.tables[tbl.Index]

	if stmt.ValuesList != nil {
		/* INSERT INTO ... VALUES ... */
		plan.ExprsList, err = ana.transformInsertValues(tbl, stmt.ValuesList)
		if err != nil {
			return nil, err
		}
	} else if stmt.Select != nil {
		/* INSERT INTO ... SELECT ... */
		return nil, errlog.New("not support insert ... select ...")
	} else {
		return nil, errlog.New("not recognized insert type")
	}

	return plan, nil
}

func (ana *Analyzer) transformSelectStmt(stmt *ast.SelectStmt) (*logic.SelectPlan, error) {
	/* 生成查询树 */
	query := new(logic.QueryPlan)

	from, err := ana.transformFromClause(stmt.FromClause)
	if err != nil {
		return nil, err
	}

	query.TargetList, err = ana.transformFields(stmt.Fields)
	if err != nil {
		return nil, err
	}

	query.Qual, err = ana.transformWhereClause(stmt.WhereClause)
	if err != nil {
		return nil, err
	}

	query.Tables = ana.tables
	query.From = from

	/* 生成 select plan */
	plan := new(logic.SelectPlan)
	plan.Query = query

	plan.Limit, err = ana.transformLimitClause(stmt.LimitClause)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (ana *Analyzer) transformUpdateStmt(stmt *ast.UpdateStmt) (*logic.UpdatePlan, error) {
	/* 生成查询树 */
	query := new(logic.QueryPlan)

	tbl, err := ana.transformSpecTable(stmt.Table)
	if err != nil {
		return nil, err
	}

	query.TargetList, err = ana.transformFields([]ast.Node{&ast.FieldStar{TableName: ana.tables[tbl.Index].TableName}})
	if err != nil {
		return nil, err
	}

	query.Qual, err = ana.transformWhereClause(stmt.WhereClause)
	if err != nil {
		return nil, err
	}

	query.Tables = ana.tables
	query.From = tbl

	/* 生成 update plan */
	plan := new(logic.UpdatePlan)
	plan.Table = ana.tables[tbl.Index]
	plan.SetList, err = ana.transformSetList(tbl, stmt.SetList)
	if err != nil {
		return nil, err
	}
	plan.Query = query

	return plan, nil
}

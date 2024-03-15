package analyzer

import (
	"fmt"
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/analyzer/logic"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/parser/ast"
)

func (ana *Analyzer) transformFromClause(fromClause ast.Node) (logic.FromObject, error) {
	if fromClause == nil {
		return nil, nil
	}

	return ana.transformFromClauseItem(fromClause)
}

func (ana *Analyzer) transformFromClauseItem(n ast.Node) (logic.FromObject, error) {
	switch n := n.(type) {
	case *ast.SpecTable:
		return ana.transformSpecTable(n)
	case *ast.JoinTable:
		return ana.transformJoinTable(n)
	default:
		return nil, errlog.New(fmt.Sprintf("transformFromClauseItem not support Node type %T", n))
	}
}

func (ana *Analyzer) transformWhereClause(node ast.ExprNode) (expression.Expression, error) {
	if node == nil {
		return nil, nil
	}

	qual, err := ana.transformExprNode(node)
	if err != nil {
		return nil, err
	}

	return ana.castToTargetType(qual, metadata.BooleanID)
}

func (ana *Analyzer) transformLimitClause(node *ast.LimitClause) (*logic.LimitPlan, error) {
	if node == nil {
		return nil, nil
	}

	limit := new(logic.LimitPlan)

	/* 处理 count 相关信息 */
	if node.Count != nil {
		count, err := ana.transformExprNode(node.Count)
		if err != nil {
			return nil, err
		}
		count, err = ana.castToTargetType(count, metadata.Uint64ID)
		if err != nil {
			return nil, err
		}
		limit.Count = count
	}

	/* 处理 offset 相关信息 */
	if node.Offset != nil {
		offset, err := ana.transformExprNode(node.Offset)
		if err != nil {
			return nil, err
		}
		offset, err = ana.castToTargetType(offset, metadata.Uint64ID)
		if err != nil {
			return nil, err
		}
		limit.Offset = offset
	}

	return limit, nil
}

package row

import (
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/mana/datum"
)

/*
 * 执行投影操作, 输出列的数据
 */
func ExecProject(project []expression.Expression, outer, inner []datum.Datum) ([]datum.Datum, error) {
	values := make([]datum.Datum, len(project))
	for i, expr := range project {
		value, err := expr.Evaluate(outer, inner)
		if err != nil {
			return nil, err
		}
		values[i] = value
	}
	return values, nil
}

func ExecQual(qual expression.Expression, outer, inner []datum.Datum) (bool, error) {
	res, err := qual.Evaluate(outer, inner)
	if err != nil || datum.IsNull(res) {
		return false, err
	}

	return datum.DatumGetValue[bool](res), nil
}

package expression

import (
	"fmt"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/mana/metadata"
)

/*
 * 变量, 表示表中的某一个列.
 */
type Variable struct {
	VarNo    int /* 所在表的序号 */
	VarColNo int /* 所在表中的列号 */
	VarType  metadata.OID
}

const (
	OuterVar int = -1
	InnerVar int = -2
)

func (v *Variable) TypeID() metadata.OID {
	return v.VarType
}

func (v *Variable) Walker(walk Walker) {

}

func (v *Variable) Mutator(mutate Mutator) Expression {
	result := *v
	return &result
}

func (v *Variable) Evaluate(outer, inner []datum.Datum) (datum.Datum, error) {
	switch v.VarNo {
	case OuterVar:
		return outer[v.VarColNo], nil
	case InnerVar:
		return inner[v.VarColNo], nil
	default:
		return nil, errlog.New("var no error ")
	}
}

func (v *Variable) ToString() (string, error) {
	return fmt.Sprintf("$%d.%d", -(v.VarNo + 1), v.VarColNo), nil
}

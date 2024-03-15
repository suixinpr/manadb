package common

import (
	"github/suixinpr/manadb/internal/mana/metadata"
)

type EntryDesc struct {
	Cols       []*metadata.ManaColumns
	Constraint any
}

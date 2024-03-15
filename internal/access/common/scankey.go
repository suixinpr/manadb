package common

import (
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
)

type ScanKeyStrategy uint16

const (
	SKInvalidStrategy ScanKeyStrategy = iota
	SKLessStrategy
	SKLessEqualStrategy
	SKEqualStrategy
	SKGreaterEqualStrategy
	SKGreaterStrategy
)

type ScanKey struct {
	ColNo     int16
	Strategy  ScanKeyStrategy
	Argument  datum.Datum
	CmpFuncID metadata.OID
}

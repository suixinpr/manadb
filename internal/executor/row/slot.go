package row

import (
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/storage/page"
)

/* 表中一行的值 */
type EntrySlot struct {
	/* entry 的值 */
	Values []datum.Datum

	/* entry 的位置, 如果不是真实存在的 entry 则为 nil */
	Pos *page.EntryPos
}

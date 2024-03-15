package session

import (
	"bufio"
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/executor"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/relation/catalog"
	"github/suixinpr/manadb/internal/utils/fmngr"
	"github/suixinpr/manadb/pkg/protocol"
)

func sendDescription(output *bufio.Writer, desc *common.EntryDesc) error {
	protocol.SendMessage(output, protocol.MessageRowDesc)
	cols := desc.Cols
	protocol.SendInt16(output, int16(len(cols)))
	for _, col := range cols {
		protocol.SendUint64(output, uint64(col.ColTypeID))
		protocol.SendString(output, col.ColName)
	}
	return nil
}

func sendResultSQL(output *bufio.Writer, desc *common.EntryDesc, ch <-chan *executor.Result) error {
	for res := range ch {
		if res.Err != nil {
			return res.Err
		}

		protocol.SendMessage(output, protocol.MessageRowData)
		slot := res.Slot
		for i, value := range slot.Values {
			if value == nil {
				protocol.SendInt32(output, -1)
				continue
			}

			col := desc.Cols[i]
			meta, _ := catalog.SearchCatalogOne(metadata.ManaTypesID, []*common.ScanKey{
				catalog.NewScanKey(metadata.ManaTypesTypeID, metadata.OIDID, datum.ValueGetDatum(col.ColTypeID)),
			})
			res, _ := fmngr.Call(datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaTypesTypeOutput]), value)
			protocol.SendDatum(output, res)
		}
	}
	return nil
}

func sendError(output *bufio.Writer, err error) error {
	protocol.SendMessage(output, protocol.MessageError)
	protocol.SendString(output, err.Error())
	return output.Flush()
}

package page

import (
	"bytes"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"testing"
)

func TestPage(t *testing.T) {
	test := []struct {
		def    []*metadata.ManaColumns
		rows   [][]datum.Datum
		header *PageHeader
	}{
		{
			def: []*metadata.ManaColumns{
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.BooleanID, ColTypeLen: 1, ColTypeAlign: 1, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.Float32ID, ColTypeLen: 4, ColTypeAlign: 4, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.Float64ID, ColTypeLen: 8, ColTypeAlign: 8, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.Int8ID, ColTypeLen: 1, ColTypeAlign: 1, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.Int16ID, ColTypeLen: 2, ColTypeAlign: 2, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.Int32ID, ColTypeLen: 4, ColTypeAlign: 4, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.Int64ID, ColTypeLen: 8, ColTypeAlign: 8, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.Uint8ID, ColTypeLen: 1, ColTypeAlign: 1, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.Uint16ID, ColTypeLen: 2, ColTypeAlign: 2, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.Uint32ID, ColTypeLen: 4, ColTypeAlign: 4, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.Uint64ID, ColTypeLen: 8, ColTypeAlign: 8, ColNotNULL: false},
			},
			rows: [][]datum.Datum{
				{
					datum.ValueGetDatum[bool](true),
					datum.ValueGetDatum[float32](3.14),
					datum.ValueGetDatum[float64](-3.1415926),
					datum.ValueGetDatum[int8](1),
					datum.ValueGetDatum[int16](-1 << 10),
					datum.ValueGetDatum[int32](1 << 20),
					datum.ValueGetDatum[int64](-1 << 40),
					datum.ValueGetDatum[uint8](1),
					datum.ValueGetDatum[uint16](1 << 10),
					datum.ValueGetDatum[uint32](1 << 20),
					datum.ValueGetDatum[uint64](1 << 40),
				},
				{
					datum.ValueGetDatum[bool](true),
					datum.ValueGetDatum[float32](3.14),
					datum.ValueGetDatum[float64](-3.1415926),
					datum.ValueGetDatum[int8](1),
					datum.ValueGetDatum[int16](-1 << 10),
					datum.ValueGetDatum[int32](1 << 20),
					datum.ValueGetDatum[int64](-1 << 40),
					datum.ValueGetDatum[uint8](1),
					datum.ValueGetDatum[uint16](1 << 10),
					datum.ValueGetDatum[uint32](1 << 20),
					datum.ValueGetDatum[uint64](1 << 40),
				},
				{
					datum.ValueGetDatum[bool](true),
					datum.ValueGetDatum[float32](3.14),
					datum.ValueGetDatum[float64](-3.1415926),
					datum.ValueGetDatum[int8](1),
					datum.ValueGetDatum[int16](-1 << 10),
					datum.ValueGetDatum[int32](1 << 20),
					datum.ValueGetDatum[int64](-1 << 40),
					datum.ValueGetDatum[uint8](1),
					datum.ValueGetDatum[uint16](1 << 10),
					datum.ValueGetDatum[uint32](1 << 20),
					datum.ValueGetDatum[uint64](1 << 40),
				},
			},
			header: &PageHeader{
				Checksum: 2687097995276241877,
				Lower:    28,
				Upper:    8024,
				Special:  8192,
				Flags:    0,
			},
		},
		{
			def: []*metadata.ManaColumns{
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.CharID, ColTypeLen: -1, ColTypeAlign: 1, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.CharID, ColTypeLen: -1, ColTypeAlign: 1, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.TextID, ColTypeLen: -1, ColTypeAlign: 1, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.VarcharID, ColTypeLen: -1, ColTypeAlign: 1, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.VarcharID, ColTypeLen: -1, ColTypeAlign: 1, ColNotNULL: false},
			},
			rows: [][]datum.Datum{
				{
					datum.StringGetDatum("abcd"),
					datum.StringGetDatum("abcdefghijklmnop"),
					datum.StringGetDatum("xxxxxxxxxx"),
					datum.StringGetDatum("01"),
					datum.StringGetDatum("0123456789"),
				},
				{
					datum.StringGetDatum("abcd"),
					datum.StringGetDatum("abcdefghijklmnop"),
					datum.StringGetDatum("xxxxxxxxxx"),
					datum.StringGetDatum("01"),
					datum.StringGetDatum("0123456789"),
				},
				{
					datum.StringGetDatum("abcd"),
					datum.StringGetDatum("abcdefghijklmnop"),
					datum.StringGetDatum("xxxxxxxxxx"),
					datum.StringGetDatum("01"),
					datum.StringGetDatum("0123456789"),
				},
			},
			header: &PageHeader{
				Checksum: 2687097995276241877,
				Lower:    28,
				Upper:    8000,
				Special:  8192,
				Flags:    0,
			},
		},
		{
			def: []*metadata.ManaColumns{
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.OIDID, ColTypeLen: 8, ColTypeAlign: 8, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.OIDID, ColTypeLen: 8, ColTypeAlign: 8, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.OIDID, ColTypeLen: 8, ColTypeAlign: 8, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.OIDID, ColTypeLen: 8, ColTypeAlign: 8, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.OIDArrayID, ColTypeLen: -1, ColTypeAlign: 8, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.OIDArrayID, ColTypeLen: -1, ColTypeAlign: 8, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.OIDArrayID, ColTypeLen: -1, ColTypeAlign: 8, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.OIDArrayID, ColTypeLen: -1, ColTypeAlign: 8, ColNotNULL: false},
			},
			rows: [][]datum.Datum{
				{
					nil,
					nil,
					datum.ValueGetDatum[metadata.OID](9999),
					datum.ValueGetDatum[metadata.OID](0),
					datum.ValueArrayGetDatum[metadata.OID]([]metadata.OID{1, 2, 3, 4}),
					datum.ValueArrayGetDatum[metadata.OID]([]metadata.OID{1234, 999, 0, 1000000}),
					nil,
					nil,
				},
				{
					nil,
					nil,
					datum.ValueGetDatum[metadata.OID](9999),
					datum.ValueGetDatum[metadata.OID](0),
					datum.ValueArrayGetDatum[metadata.OID]([]metadata.OID{1, 2, 3, 4}),
					datum.ValueArrayGetDatum[metadata.OID]([]metadata.OID{1234, 999, 0, 1000000}),
					nil,
					nil,
				},
				{
					nil,
					nil,
					datum.ValueGetDatum[metadata.OID](9999),
					datum.ValueGetDatum[metadata.OID](0),
					datum.ValueArrayGetDatum[metadata.OID]([]metadata.OID{1, 2, 3, 4}),
					datum.ValueArrayGetDatum[metadata.OID]([]metadata.OID{1234, 999, 0, 1000000}),
					nil,
					nil,
				},
			},
			header: &PageHeader{
				Checksum: 2687097995276241877,
				Lower:    28,
				Upper:    7880,
				Special:  8192,
				Flags:    0,
			},
		},
	}
	for no, tt := range test {
		p := NewPage()
		p.Init()
		for _, values := range tt.rows {
			e := NewEntry(tt.def, values)
			p.InsertEntry(e)
		}

		if *tt.header != *p.GetHeader() {
			t.Errorf("(%v) GetHeader: expected %v, got = %v", no, *tt.header, *p.GetHeader())
		}

		max := p.GetMaxItemID()
		for itemID := OffsetNumber(0); itemID < max; itemID++ {
			item := p.GetItem(itemID)
			e := p.GetEntry(item)
			values := e.GetAllColumnsValue(tt.def)

			num := len(values)
			for i := 0; i < num; i++ {
				if !bytes.Equal(tt.rows[max-itemID-1][i], values[i]) {
					t.Errorf("(%v) GetAllColumnsValue: expected %v, got = %v", no, tt.rows[max-itemID-1][i], values[i])
				}
			}
		}
	}
}

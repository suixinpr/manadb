package page

import (
	"bytes"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"testing"
)

func TestEntry(t *testing.T) {
	test := []struct {
		def    []*metadata.ManaColumns
		values []datum.Datum
		entry  Entry
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
			values: []datum.Datum{
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
			entry: Entry{11, 0, 0, 0, 1, 0, 0, 0, 195, 245, 72, 64, 0, 0, 0, 0, 74, 216, 18, 77, 251, 33, 9, 192, 1, 0, 0, 252, 0, 0, 16, 0, 0, 0, 0, 0, 0, 255, 255, 255, 1, 0, 0, 4, 0, 0, 16, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		},
		{
			def: []*metadata.ManaColumns{
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.CharID, ColTypeLen: -1, ColTypeAlign: 1, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.CharID, ColTypeLen: -1, ColTypeAlign: 1, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.TextID, ColTypeLen: -1, ColTypeAlign: 1, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.VarcharID, ColTypeLen: -1, ColTypeAlign: 1, ColNotNULL: false},
				{ColID: 0, ColTblID: 0, ColName: "", ColNo: 0, ColMod: 0, ColTypeID: metadata.VarcharID, ColTypeLen: -1, ColTypeAlign: 1, ColNotNULL: false},
			},
			values: []datum.Datum{
				datum.StringGetDatum("abcd"),
				datum.StringGetDatum("abcdefghijklmnop"),
				datum.StringGetDatum("xxxxxxxxxx"),
				datum.StringGetDatum("01"),
				datum.StringGetDatum("0123456789"),
			},
			entry: Entry{5, 0, 0, 0, 4, 0, 97, 98, 99, 100, 16, 0, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 10, 0, 120, 120, 120, 120, 120, 120, 120, 120, 120, 120, 2, 0, 48, 49, 10, 0, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 0, 0, 0, 0, 0, 0, 0, 0},
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
			values: []datum.Datum{
				nil,
				nil,
				datum.ValueGetDatum[metadata.OID](9999),
				datum.ValueGetDatum[metadata.OID](0),
				datum.ValueArrayGetDatum[metadata.OID]([]metadata.OID{1, 2, 3, 4}),
				datum.ValueArrayGetDatum[metadata.OID]([]metadata.OID{1234, 999, 0, 1000000}),
				nil,
				nil,
			},
			entry: Entry{8, 0, 1, 0, 195, 0, 0, 0, 15, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 32, 0, 0, 0, 0, 0, 0, 0, 210, 4, 0, 0, 0, 0, 0, 0, 231, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 66, 15, 0, 0, 0, 0, 0},
		},
	}
	for no, tt := range test {
		num := len(tt.def)
		e := NewEntry(tt.def, tt.values)
		if !bytes.Equal(tt.entry, e) {
			t.Errorf("(%v) NewEntry: expected %v, got = %v", no, tt.entry, e)
		}
		for i := 0; i < num; i++ {
			res := e.GetColumnValue(tt.def, int16(i))
			if !bytes.Equal(tt.values[i], res) {
				t.Errorf("(%v) GetColumnValue: values[%v] expected %v, got = %v", no, i, tt.values[i], res)
			}
		}
		res := e.GetAllColumnsValue(tt.def)
		for i := 0; i < num; i++ {
			if !bytes.Equal(tt.values[i], res[i]) {
				t.Errorf("(%v) GetAllColumnsValue: values[%v] expected %v, got = %v", no, i, tt.values[i], res[i])
			}
		}
	}
}

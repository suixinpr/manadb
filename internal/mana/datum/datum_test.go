package datum

import (
	"fmt"
	"testing"
)

func TestNull(t *testing.T) {
	s1 := ""
	d1 := StringGetDatum(s1)
	fmt.Println(d1 == nil)

	var b []byte
	fmt.Println(b == nil)

	b2 := BytesGetDatum(b)
	fmt.Println(b2 == nil)

	fmt.Println(EmtryDatum() == nil, NullDatum() == nil)

	fmt.Println(DatumGetBytes(nil) == nil)

	fmt.Println(DatumGetString(nil) == "")

	p := []byte{'1', '2'}
	fmt.Println(Datum(p[0:0]) == nil)

	t.Error("")
}

package buffer

/*
import (
	"github/suixinpr/manadb/internal/storage/smngr"
	"strconv"
	"testing"
)

func TestNewBufferPool(t *testing.T) {
	test := []struct {
		name string

		capacity  uint64
		bucketNum uint64
	}{
		{"DefaultBufferPool", 2048, 256},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			bmgr, err := NewBufferPool(tt.capacity, tt.bucketNum, 1<<10)
			if err != nil {
				t.Errorf("NewBufferPool() err: %v", err)
			}
			if len(bmgr.bufferPool) != int(tt.capacity) {
				t.Errorf("NewBufferPool() capacity: got = %v, want = %v", len(bmgr.bufferPool), int(tt.capacity))
			}
			if len(bmgr.bufferMap) != int(tt.bucketNum) {
				t.Errorf("NewBufferPool() bucketNum: got = %v, want = %v", len(bmgr.bufferMap), int(tt.bucketNum))
			}
		})
	}
}

func TestGetBuffer(t *testing.T) {
	test := []struct {
		name string

		pi smngr.PageIdentifier
	}{
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
	}

	bmgr, err := NewBufferPool(4, 2, 1<<10)
	if err != nil {
		t.Errorf("TestGetBuffer() NewBufferPool err: %v", err)
		return
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			buf, err := bmgr.GetBuffer(tt.pi)
			if err != nil {
				t.Errorf("TestGetBuffer() err: %v", err)
			}
			if buf.pi != tt.pi {
				t.Errorf("TestGetBuffer() pi: got = %v, want = %v", buf.pi, tt.pi)
			}
			buf.Release()
		})
	}
}

func TestParallelGetBuffer(t *testing.T) {
	processNum := 100000
	test := []struct {
		name string

		pi smngr.PageIdentifier
	}{
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
	}

	bmgr, err := NewBufferPool(4, 2, 1<<10)
	if err != nil {
		t.Errorf("TestGetBuffer() NewBufferPool err: %v", err)
		return
	}

	for i := 0; i < processNum; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			for _, tt := range test {
				buf, err := bmgr.GetBuffer(tt.pi)
				if err != nil {
					t.Errorf("TestGetBuffer() err: %v", err)
				}
				if buf.pi != tt.pi {
					t.Errorf("TestGetBuffer() pi: got = %v, want = %v", buf.pi, tt.pi)
				}
				buf.Release()
			}
		})
	}
}
*/

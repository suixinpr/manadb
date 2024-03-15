package smngr

import (
	"testing"
)

func TestStorageManager(t *testing.T) {
	var err error
	smgr := OpenStorageManager(0)

	err = smgr.Create()
	if err != nil {
		t.Error(err)
	}

	n, err := smgr.PageNum()
	if err != nil {
		t.Error(err)
	}
	if n != 0 {
		t.Error(n)
	}
	if err := smgr.Extend(); err != nil {
		t.Error(err)
	}
	n, err = smgr.PageNum()
	if err != nil {
		t.Error(err)
	}
	if n != 1 {
		t.Error(n)
	}
	if err := smgr.Extend(); err != nil {
		t.Error(err)
	}
	n, err = smgr.PageNum()
	if err != nil {
		t.Error(err)
	}
	if n != 2 {
		t.Error(n)
	}

	if err := smgr.Close(); err != nil {
		t.Error(err)
	}

	if err := smgr.Remove(); err != nil {
		t.Error(err)
	}
}

package errlog

import (
	"errors"
	"runtime/debug"
)

func New(s string) error {
	err := errors.New(s)
	debug.PrintStack()
	return err
}

func Err(err error) error {
	if err == nil {
		return nil
	}
	debug.PrintStack()
	return err
}

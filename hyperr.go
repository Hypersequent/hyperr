package hyperr

import (
	"fmt"
	"path"
	"runtime"
	"strconv"
)

type Error struct {
	Err    error
	Ann    string
	Caller string
}

func (e Error) Unwrap() error {
	return e.Err
}

func (e Error) Error() string {
	if e.Ann == "" {
		return fmt.Sprintf("[%s]: %s", e.Caller, e.Err.Error())
	}
	return fmt.Sprintf("[%s] %s: %s", e.Caller, e.Ann, e.Err.Error())
}

func Wrap(err error) Error {
	return Error{
		Err:    err,
		Caller: caller(2),
	}
}

func Annotate(err error, f string, args ...interface{}) Error {
	return Error{
		Err:    err,
		Caller: caller(2),
		Ann:    fmt.Sprintf(f, args...),
	}
}

func caller(skip int) string {
	_, f, line, ok := runtime.Caller(skip)
	if !ok {
		return "<unknown>"
	}
	_, file := path.Split(f)

	return file + ":" + strconv.Itoa(line)
}

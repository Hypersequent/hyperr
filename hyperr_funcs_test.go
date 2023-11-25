package hyperr

import "fmt"

func wrapMe() error {
	return Wrap(fmt.Errorf("hello"))
}

func annotateMe() error {
	return Annotate(fmt.Errorf("hello"), "world")
}

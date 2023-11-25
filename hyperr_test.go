package hyperr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	assert.Equal(t, "[hyperr_funcs_test.go:6]: hello", wrapMe().Error())
}

func TestAnnotate(t *testing.T) {
	assert.Equal(t, "[hyperr_funcs_test.go:10] world: hello", annotateMe().Error())
}

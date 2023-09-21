package etrier

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTry(t *testing.T) {
	SetHandler(func(err error) {
		assert.EqualError(t, err, "strconv.Atoi: parsing \"foo\": invalid syntax")
	})
	foo(t)
}

func foo(t *testing.T) {
	t.Log(Try2(strconv.Atoi("foo")))
	
	bar(t)
}

func bar(t *testing.T) {
	SetHandler(func(err error) {
		assert.EqualError(t, err, "strconv.Atoi: parsing \"bar\": invalid syntax")
	})
	t.Log(Try2(strconv.Atoi("bar")))

	bazz(t)
}

func bazz(t *testing.T) {
	t.Log(Try2(strconv.Atoi("3")))
}

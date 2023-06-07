package encode

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestEncode(t *testing.T) {
	res := Encode("https://www.rymoe.com", "Hello World!")
	assert.Equal(t, res, "lTUcFU2hs5BWowOj")
}

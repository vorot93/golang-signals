package signals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignals(t *testing.T) {
	type TestType struct {
		MySignal Signal
	}
	var fixture TestType
	var result string
	fixture.MySignal.Connect(func(data interface{}) {
		var str, _ = data.(string)
		result = str
	})

	var expectation = "edited"
	fixture.MySignal.Emit("edited")

	assert.Equal(t, expectation, result)
}

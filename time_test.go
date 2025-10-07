package debugo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	assert.Equal(t, "1m0s0ms", prettyPrintDuration(1*time.Minute))
	assert.Equal(t, "1m30s0ms", prettyPrintDuration(90*time.Second))
	assert.Equal(t, "1m30s500ms", prettyPrintDuration(90*time.Second+500*time.Millisecond))
	assert.Equal(t, "1h1m30s500ms", prettyPrintDuration(90*time.Second+500*time.Millisecond+1*time.Hour))
}

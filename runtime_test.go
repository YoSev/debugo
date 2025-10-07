package debugo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	SetOutput(os.Stderr)
	assert.Equal(t, os.Stderr, GetOutput())
	SetOutput(os.Stdout)
	assert.Equal(t, os.Stdout, GetOutput())
	SetOutput(os.Stdin)
	assert.Equal(t, os.Stdin, GetOutput())
}

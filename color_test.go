package debugo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomColor(t *testing.T) {
	var namespace = "test"
	color := getRandomColor(namespace)
	assert.NotNil(t, color, "Expect color to exist")
	assert.Equal(t, color, getRandomColor(namespace), "Expect colors to match")
}

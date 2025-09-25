package debugo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	d := New("namespace")
	assert.Equal(t, "namespace", d.namespace, "Check new debugger")

	d = d.Extend("moo")
	assert.Equal(t, "namespace:moo", d.namespace, "Check new debugger")
}

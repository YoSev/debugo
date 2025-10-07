package debugo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchNamespace(t *testing.T) {
	l := New("test:a")

	// Empty debug string
	SetNamespace("")
	assert.False(t, l.matchNamespace(), "they should be equal")

	// Single wildcard
	SetNamespace("*")
	assert.True(t, l.matchNamespace(), "they should be equal")

	// Single negative pattern
	SetNamespace("-test:*")
	assert.False(t, l.matchNamespace(), "they should be equal")

	// Multiple negative patterns
	SetNamespace("-test:a,-test:b")
	assert.False(t, l.matchNamespace(), "they should be equal")

	// Leading/trailing spaces in debug string
	SetNamespace("   test:a   ")
	assert.True(t, l.matchNamespace(), "they should be equal")

	// Overlapping inclusion and exclusion
	SetNamespace("test:*, -test:a")
	assert.False(t, l.matchNamespace(), "they should be equal")

	// Overlapping inclusion and exclusion
	SetNamespace("*, -test:a")
	assert.False(t, l.matchNamespace(), "they should be equal")

	// Overlapping inclusion and exclusion
	SetNamespace("*, -test:?")
	assert.False(t, l.matchNamespace(), "they should be equal")

	// Overlapping inclusion and exclusion
	SetNamespace("*, -test:*")
	assert.False(t, l.matchNamespace(), "they should be equal")

	// Exact match
	SetNamespace("test:a")
	assert.True(t, l.matchNamespace(), "they should be equal")

	// Case sensitivity
	SetNamespace("TEST:A")
	assert.True(t, l.matchNamespace(), "they should be equal") // Assuming case-sensitive

	// No namespace provided
	l = New("")
	SetNamespace("test:*")
	assert.False(t, l.matchNamespace(), "they should be equal")

	// Non-matching wildcards
	SetNamespace("test:x*")
	assert.False(t, l.matchNamespace(), "they should be equal")

	// Invalid patterns
	SetNamespace("--test:*")
	assert.False(t, l.matchNamespace(), "they should be equal")

	// --- Optional pattern tests ---
	l = New("info")
	SetNamespace("info:?")
	assert.True(t, l.matchNamespace(), "should match the base namespace")

	l = New("info:test")
	SetNamespace("info:?")
	assert.True(t, l.matchNamespace(), "should match the namespace with suffix")

	l = New("info:test:sub")
	SetNamespace("info:?")
	assert.True(t, l.matchNamespace(), "should match namespace with deeper suffix")

	l = New("other")
	SetNamespace("info:?")
	assert.False(t, l.matchNamespace(), "should not match unrelated namespace")
}

func TestMatchPattern(t *testing.T) {
	tests := []struct {
		namespace string
		pattern   string
		expected  bool
	}{
		// Exact match
		{"test:a", "test:a", true},
		{"test:a", "test:b", false},

		// Wildcard match
		{"test:a", "test:*", true},    // Any namespace starting with "test:"
		{"test:abc", "test:*", true},  // Any namespace starting with "test:"
		{"abc:test", "test:*", false}, // Doesn't match if the prefix is different

		// Wildcard with empty prefix
		{"test:a", "*", true}, // "*" should match anything
		{"", "*", true},       // "*" matches empty string

		// Wildcard in middle
		{"test:foo:bar", "test:*:bar", true}, // Should match, as it fits the pattern

		// Mismatch cases
		{"test:a", "test:b", false},  // Doesn't match if the pattern is different
		{"test:a", "other:*", false}, // Doesn't match if the prefix is different

		// Pattern with trailing `*`
		{"test:abc", "test:*", true}, // Should match any string starting with "test:"
		{"test:xyz", "test:*", true}, // Should match any string starting with "test:"

		// Multiple wildcard checks
		{"test:abc", "test:*:*", false},   // Matches with `test:*:*`
		{"test:abc", "test:abc:*", false}, // Matches with `test:abc:*`
		{"test:abc", "test:*:abc", false}, // Doesn't match "test:abc" with "test:*:abc"

		// --- Optional pattern tests ---
		{"info", "info:?", true},          // exact base match
		{"info:test", "info:?", true},     // base with suffix
		{"info:test:sub", "info:?", true}, // deeper suffix
		{"other", "info:?", false},        // non-matching
	}

	for _, tt := range tests {
		t.Run(tt.namespace+"_"+tt.pattern, func(t *testing.T) {
			result := matchPattern(tt.namespace, tt.pattern)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMatchpatterns(t *testing.T) {
	ok := matchPattern("namespace", "[invalid")
	assert.False(t, ok)
	ok = matchPattern("namespace", "[invalid:?")
	assert.False(t, ok)
}

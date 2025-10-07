package debugo

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var namespace = "test-namespace"

var testMessage = "Test message"
var testMessageExpected = fmt.Sprintf("%s %s +0ms\n", namespace, testMessage)

var testMessageWithFormatting = "Test message with %d"
var testMessageWithFormattingArgs = 42
var testMessageWithFormattingExpected = fmt.Sprintf("%s Test message with %d +0ms\n", namespace, testMessageWithFormattingArgs)

func getDebugger() *Debugger {
	SetNamespace("*")
	return New(namespace)
}

func stripANSI(input string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(input, "")
}

func hasANSI(input string) bool {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.MatchString(input)
}

func TestDebug(t *testing.T) {
	var buf bytes.Buffer
	d := getDebugger()
	d.SetOutput(&buf)

	d.Debug(testMessage)

	assert.True(t, hasANSI(buf.String()))

	output := strings.TrimSpace(stripANSI(buf.String())) // Strip colors and trim whitespace
	expected := strings.TrimSpace(testMessageExpected)
	assert.Equal(t, expected, output)
}

func TestDebugJSON(t *testing.T) {
	var buf bytes.Buffer
	d := getDebugger()
	d.SetOutput(&buf)
	SetFormat(JSON)
	d.Debug(testMessage)
	SetFormat(Plain)

	assert.False(t, hasANSI(buf.String()))
	output := strings.TrimSpace(stripANSI(buf.String())) // Strip colors and trim whitespace
	expected := strings.TrimSpace("{\"namespace\":\"" + namespace + "\",\"message\":\"" + testMessage + "\"}")
	assert.Equal(t, expected, output)
}

func TestDebugJSONWithTimestamp(t *testing.T) {
	var buf bytes.Buffer
	d := getDebugger()
	SetOutput(&buf)
	SetTimestamp(&Timestamp{Format: "2006"})
	SetFormat(JSON)
	d.Debug(testMessage)
	SetFormat(Plain)
	SetTimestamp(nil)

	assert.False(t, hasANSI(buf.String()))
	output := strings.TrimSpace(stripANSI(buf.String())) // Strip colors and trim whitespace
	expected := strings.TrimSpace("{\"timestamp\":\"" + fmt.Sprint(time.Now().Year()) + "\",\"namespace\":\"" + namespace + "\",\"message\":\"" + testMessage + "\"}")
	assert.Equal(t, expected, output)
}

func TestDebugGlobalOutput(t *testing.T) {
	var buf bytes.Buffer
	d := getDebugger()
	d.SetOutput(&buf)
	SetOutput(&buf)
	d.SetOutput(nil)

	d.Debug(testMessage)

	assert.True(t, hasANSI(buf.String()))

	output := strings.TrimSpace(stripANSI(buf.String())) // Strip colors and trim whitespace
	expected := strings.TrimSpace(testMessageExpected)
	assert.Equal(t, expected, output)
}

func TestDebugNoColors(t *testing.T) {
	var buf bytes.Buffer
	SetUseColors(false)
	d := getDebugger()
	d.SetOutput(&buf)

	d.Debug(testMessage)

	assert.False(t, hasANSI(buf.String()))
}

func TestDebugNonMatchingNamespace(t *testing.T) {
	var buf bytes.Buffer
	SetUseColors(false)
	d := getDebugger()
	d.SetOutput(&buf)

	d.Debug("")

	assert.Empty(t, buf.String())
}

func TestDebugEmptyMessage(t *testing.T) {
	var buf bytes.Buffer
	SetUseColors(false)
	d := getDebugger()
	d.SetOutput(&buf)

	SetNamespace("does:not:exist")
	d.Debug("test")

	assert.Empty(t, buf.String())
}

func TestDebugWithColors(t *testing.T) {
	var buf bytes.Buffer
	SetUseColors(true)
	d := getDebugger()
	d.SetOutput(&buf)

	d.Debug(testMessage)

	assert.True(t, hasANSI(buf.String()))
}

func TestDebugf(t *testing.T) {
	var buf bytes.Buffer
	d := getDebugger()
	d.SetOutput(&buf)

	d.Debugf(testMessageWithFormatting, testMessageWithFormattingArgs)

	output := strings.TrimSpace(stripANSI(buf.String())) // Strip colors and trim whitespace
	expected := strings.TrimSpace(testMessageWithFormattingExpected)
	if output != expected {
		t.Errorf("Expected '%s' in output, got: '%s'", expected, output)
	}
}

func TestDebugRaceCondition(_ *testing.T) {
	var buf bytes.Buffer
	d := getDebugger()
	d.SetOutput(&buf)

	const goroutines = 10
	const iterations = 100
	done := make(chan bool, goroutines)

	ids := make([]int, goroutines)
	for i := range ids {
		ids[i] = i
	}

	for _, id := range ids {
		go func(id int) {
			for j := 0; j < iterations; j++ {
				SetNamespace("*")
				SetTimestamp(&Timestamp{Format: time.RFC3339})
				d.Debug(fmt.Sprintf("Goroutine %d, iteration %d", id, j))
			}
			done <- true
		}(id)
	}

	for range ids {
		<-done
	}

	// Optionally, verify output without colors
	_ = stripANSI(buf.String())
}

func TestWriteWithFields(t *testing.T) {
	var buf bytes.Buffer
	d := getDebugger().With("foo", "bar").With("", "empty-key").With("foo", nil).With("func", func() {})
	d.SetOutput(&buf)
	SetFormat(Plain)
	SetUseColors(false)
	SetTimestamp(&Timestamp{Format: "2006"})
	d.Debugf("%s %s %t", "foo", "bar", true)

	t.Log(buf.String())
}

func TestJSONWritePrint(t *testing.T) {
	var buf bytes.Buffer
	d := getDebugger().With("foo", "bar").With("", "empty-key").With("foo", nil).With("func", func() {}).With("age", 42).With("is", true)
	d.SetOutput(&buf)
	SetFormat(JSON)
	SetTimestamp(&Timestamp{Format: "2006"})
	d.Debug("foo", "bar", true)
	assert.Equal(t, "{\"timestamp\":\"2025\",\"namespace\":\"test-namespace\",\"fields\":{\"(empty)\":\"empty-key\",\"age\":42,\"foo\":null,\"func\":\"(not serializable)\",\"is\":true},\"message\":\"foobartrue\"}\n", buf.String())
	t.Log(buf.String())
}

func TestPlainWritePrint(t *testing.T) {
	var buf bytes.Buffer
	d := getDebugger()
	d.SetOutput(&buf)
	SetTimestamp(&Timestamp{Format: time.Kitchen})
	SetUseColors(false)
	SetFormat(Plain)
	SetTimestamp(&Timestamp{Format: time.Kitchen})
	d.Debugf("%s %s %t", "foo", "bar", true)
	t.Log(buf.String())
}

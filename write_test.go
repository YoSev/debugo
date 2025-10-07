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
	SetFormat(Json)
	d.Debug(testMessage)
	SetFormat(Plain)

	assert.False(t, hasANSI(buf.String()))
	output := strings.TrimSpace(stripANSI(buf.String())) // Strip colors and trim whitespace
	expected := strings.TrimSpace("{\"namespace\":\"" + namespace + "\",\"message\":\"" + testMessage + "\"}")
	assert.Equal(t, expected, output)
}

func TestDebugJSONWithField(t *testing.T) {
	var buf bytes.Buffer
	d := getDebugger()
	d.SetOutput(&buf)
	d = d.With("json", true).With("number", 123)
	SetFormat(Json)
	d.Debug(testMessage)
	SetFormat(Plain)

	assert.False(t, hasANSI(buf.String()))
	output := strings.TrimSpace(stripANSI(buf.String())) // Strip colors and trim whitespace
	expected := strings.TrimSpace("{\"namespace\":\"" + namespace + "\",\"fields\":{\"json\":true,\"number\":123},\"message\":\"" + testMessage + "\"}")
	assert.Equal(t, expected, output)
}

func TestDebugJSONWithFieldTimestamp(t *testing.T) {
	var buf bytes.Buffer
	d := getDebugger()
	d.SetOutput(&buf)

	d = d.With("json", true).With("number", 123)
	SetTimestamp(&Timestamp{Format: "2006"})
	SetFormat(Json)
	d.Debug(testMessage)
	SetFormat(Plain)
	SetTimestamp(nil)

	assert.False(t, hasANSI(buf.String()))
	output := strings.TrimSpace(stripANSI(buf.String())) // Strip colors and trim whitespace
	expected := strings.TrimSpace("{\"timestamp\":\"" + fmt.Sprint(time.Now().Year()) + "\",\"namespace\":\"" + namespace + "\",\"fields\":{\"json\":true,\"number\":123},\"message\":\"" + testMessage + "\"}")
	assert.Equal(t, expected, output)
}

func TestDebugWithFields(t *testing.T) {
	var buf bytes.Buffer
	d := getDebugger()
	d.SetOutput(&buf)

	x := d.With("key1", "value1").With("key2", 42)

	x.Debug(testMessage)

	assert.True(t, hasANSI(buf.String()))

	output := strings.TrimSpace(stripANSI(buf.String())) // Strip colors and trim whitespace
	expected := strings.TrimSpace(fmt.Sprintf("%s key1=value1 key2=42 %s +0ms\n", namespace, testMessage))
	assert.Equal(t, expected, output, "Must have fields")
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

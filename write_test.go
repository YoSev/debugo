package debugo

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"
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

func TestDebug(t *testing.T) {
	var buf bytes.Buffer
	d := getDebugger()
	d.SetOutput(&buf)

	d.Debug(testMessage)

	output := strings.TrimSpace(stripANSI(buf.String())) // Strip colors and trim whitespace
	expected := strings.TrimSpace(testMessageExpected)
	if output != expected {
		t.Errorf("Expected '%s' in output, got: '%s'", expected, output)
	}
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

func TestDebugRaceCondition(t *testing.T) {
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

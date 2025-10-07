package debugo

import (
	"io"
	"os"
	"sync"
)

type config struct {
	namespace string
	timestamp *Timestamp

	output io.Writer

	useColors bool

	mutex *sync.RWMutex
}

var runtime = &config{
	namespace: "*",
	timestamp: nil,

	output: os.Stderr,

	useColors: true,

	mutex: &sync.RWMutex{},
}

type Timestamp struct {
	Format string
}

// SetUseColors sets the global color usage for debugging.
func SetUseColors(use bool) {
	runtime.mutex.Lock()
	defer runtime.mutex.Unlock()

	runtime.useColors = use
}

// GetUseColors retrieves the current global color usage for debugging.
func GetUseColors() bool {
	runtime.mutex.RLock()
	defer runtime.mutex.RUnlock()

	return runtime.useColors
}

// SetNamespace sets the global namespace for debugging.
func SetNamespace(namespace string) {
	runtime.mutex.Lock()
	defer runtime.mutex.Unlock()

	runtime.namespace = namespace
}

// GetNamespace retrieves the current global namespace for debugging.
func GetNamespace() string {
	runtime.mutex.RLock()
	defer runtime.mutex.RUnlock()

	return runtime.namespace
}

// SetTimestamp sets the global timestamp configuration for debugging.
func SetTimestamp(timestamp *Timestamp) {
	runtime.mutex.Lock()
	defer runtime.mutex.Unlock()

	runtime.timestamp = timestamp
}

// GetTimestamp retrieves the current global timestamp configuration for debugging.
func GetTimestamp() *Timestamp {
	runtime.mutex.RLock()
	defer runtime.mutex.RUnlock()

	return runtime.timestamp
}

// SetOutput sets the global output configuration for debugging.
func SetOutput(output io.Writer) {
	runtime.mutex.Lock()
	defer runtime.mutex.Unlock()

	runtime.output = output
}

// GetOutput retrieves the current global output configuration for debugging.
func GetOutput() io.Writer {
	runtime.mutex.RLock()
	defer runtime.mutex.RUnlock()

	return runtime.output
}

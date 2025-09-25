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

	mutex *sync.Mutex
}

var runtime = &config{
	namespace: "*",
	timestamp: nil,

	output: os.Stderr,

	mutex: &sync.Mutex{},
}

type Timestamp struct {
	Format string
}

// Set the global namespace for debugging.
func SetNamespace(namespace string) {
	runtime.mutex.Lock()
	defer runtime.mutex.Unlock()

	runtime.namespace = namespace
}

// GetNamespace retrieves the current global namespace for debugging.
func GetNamespace() string {
	runtime.mutex.Lock()
	defer runtime.mutex.Unlock()

	return runtime.namespace
}

// Sets the global timestamp configuration for debugging.
func SetTimestamp(timestamp *Timestamp) {
	runtime.mutex.Lock()
	defer runtime.mutex.Unlock()

	runtime.timestamp = timestamp
}

// Retrieves the current global timestamp configuration for debugging.
func GetTimestamp() *Timestamp {
	runtime.mutex.Lock()
	defer runtime.mutex.Unlock()

	return runtime.timestamp
}

// Sets the global output configuration for debugging.
func SetOutput(output io.Writer) {
	runtime.mutex.Lock()
	defer runtime.mutex.Unlock()

	runtime.output = output
}

// Retrieves the current global output configuration for debugging.
func GetOutput() io.Writer {
	runtime.mutex.Lock()
	defer runtime.mutex.Unlock()

	return runtime.output
}

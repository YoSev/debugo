package debugo

import (
	"encoding/json"
	"io"
	"maps"
	"sync"
	"time"

	"github.com/fatih/color"
)

func init() {
	color.NoColor = false
}

type Debugger struct {
	namespace string

	color   *color.Color
	lastLog time.Time

	output io.Writer

	fields map[string]any

	mutex *sync.Mutex
}

// New creates a new debugger instance
func New(namespace string) *Debugger {
	return newDebugger(namespace)
}

// With clones the debugger instance and adds a key-value pair to its fields (json serializeable)
func (d *Debugger) With(key string, value any) *Debugger {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	n := *d
	maps.Copy(n.fields, d.fields)

	if key == "" {
		key = "(empty)"
	}

	if value == nil {
		value = nil
	}

	if _, err := json.Marshal(value); err != nil {
		value = "(not serializable)"
	}

	n.fields[key] = value
	return &n
}

// Extend creates a new debugger instance with an extended namespace
func (d *Debugger) Extend(namespace string) *Debugger {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	n := *d
	n.namespace = d.namespace + ":" + namespace
	return &n
}

// SetOutput sets the output writer for the debugger instance
func (d *Debugger) SetOutput(output io.Writer) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.output = output
}

func newDebugger(namespace string) *Debugger {
	return &Debugger{
		namespace: namespace,

		color:   getRandomColor(namespace),
		lastLog: time.Now(),

		output: nil,

		fields: make(map[string]any),

		mutex: &sync.Mutex{}}
}

package debugo

import (
	"io"
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

	mutex *sync.Mutex
}

func New(namespace string) *Debugger {
	return newDebugger(namespace)
}

func (d *Debugger) Extend(namespace string) *Debugger {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	n := newDebugger(d.namespace + ":" + namespace)
	n.color = d.color
	n.lastLog = d.lastLog
	n.output = d.output
	n.mutex = d.mutex
	return n
}

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

		mutex: &sync.Mutex{}}
}

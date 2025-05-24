package debugo

import (
	"io"
	"os"
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
	return new(namespace)
}

func new(namespace string) *Debugger {
	return &Debugger{
		namespace: namespace,

		color:   getRandomColor(namespace),
		lastLog: time.Now(),

		output: os.Stderr,

		mutex: &sync.Mutex{}}
}

func (d *Debugger) SetOutput(output io.Writer) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if output == nil {
		output = os.Stderr
	}

	d.output = output
}

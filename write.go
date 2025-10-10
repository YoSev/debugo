package debugo

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type asJSON struct {
	Timestamp string         `json:"timestamp,omitempty"`
	Namespace string         `json:"namespace,omitempty"`
	Fields    map[string]any `json:"fields,omitempty"`
	Message   string         `json:"message,omitempty"`
	ElapsedMs int64          `json:"elapsed_ms,omitempty"`
}

// Debug logs a message if the debuggers namespace matches
func (d *Debugger) Debug(message ...any) *Debugger {
	d.write(message...)
	return d
}

// Debugf logs a formatted message if the debuggers namespace matches
func (d *Debugger) Debugf(format string, message ...any) *Debugger {
	d.write(fmt.Sprintf(format, message...))
	return d
}

func (d *Debugger) write(message ...any) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if !d.matchNamespace() {
		return
	}

	msg := fmt.Sprint(message...)
	if msg == "" {
		return
	}

	if GetFormat() == JSON {
		d.writeJSON(message...)
		return
	}
	// Optional timestamp
	var timestamp string
	if t := GetTimestamp(); t != nil {
		timestamp = time.Now().Format(t.Format)
	}

	// Build log parts
	parts := []string{}
	if timestamp != "" {
		parts = append(parts, timestamp)
	}

	if GetUseColors() {
		parts = append(parts, d.color.Sprintf("%s", d.namespace))
	} else {
		parts = append(parts, d.namespace)
	}

	// append fields as key=value
	for key, value := range d.fields {
		if value == nil {
			value = ""
		}
		parts = append(parts, fmt.Sprintf("%s=%v", key, value))
	}

	parts = append(parts, msg)
	parts = append(parts, fmt.Sprintf("+%s", prettyPrintDuration(d.elapsed())))

	log := strings.Join(parts, " ") + "\n"

	// Write to output
	if d.output != nil {
		_, _ = fmt.Fprint(d.output, log)
	} else if o := GetOutput(); o != nil {
		_, _ = fmt.Fprint(o, log)
	}
}

func (d *Debugger) writeJSON(message ...any) {
	entry := asJSON{
		Namespace: d.namespace,
		Message:   fmt.Sprint(message...),
		Fields:    d.fields,
		ElapsedMs: d.elapsed().Milliseconds(),
	}

	if t := GetTimestamp(); t != nil {
		entry.Timestamp = time.Now().Format(t.Format)
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return
	}

	if d.output != nil {
		_, _ = d.output.Write(append(data, '\n'))
	} else if o := GetOutput(); o != nil {
		_, _ = o.Write(append(data, '\n'))
	}
}

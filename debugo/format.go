package debugo

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func prettyPrintDuration(d time.Duration) string {
	// Break down the duration into hours, minutes, seconds, and milliseconds
	hours := d / time.Hour
	d %= time.Hour
	minutes := d / time.Minute
	d %= time.Minute
	seconds := d / time.Second
	d %= time.Second
	milliseconds := d / time.Millisecond

	// Build the formatted string dynamically
	result := ""
	if hours > 0 {
		result += fmt.Sprintf("%dh", hours)
	}
	if minutes > 0 || hours > 0 {
		result += fmt.Sprintf("%dm", minutes)
	}
	if seconds > 0 || minutes > 0 || hours > 0 {
		result += fmt.Sprintf("%ds", seconds)
	}
	result += fmt.Sprintf("%dms", milliseconds) // Always include milliseconds
	return result
}

func formatValue(v any) string {
	switch v := v.(type) {
	case fmt.Stringer:
		return v.String()
	default:
		rv := reflect.ValueOf(v)
		rt := reflect.TypeOf(v)

		if rv.Kind() == reflect.Struct {
			var fields []string
			for i := 0; i < rv.NumField(); i++ {
				field := rt.Field(i)
				value := rv.Field(i)
				fields = append(fields, fmt.Sprintf("%s: %v", field.Name, value.Interface()))
			}
			return fmt.Sprintf("{%s}", strings.Join(fields, ", "))
		}

		// Fallback for all other types
		return fmt.Sprintf("%v", v)
	}
}

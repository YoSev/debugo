package debugo

import (
	"os"
	"strings"
)

var debug = os.Getenv("DEBUGO")

func SetDebug(namespace string) {
	debug = namespace
}

// Enabled will return a boolean value to determine if an instance would log according to the current set debug namespace
func (l *Logger) Enabled() bool {
	return l.matchNamespace()
}

func (l *Logger) matchNamespace() bool {
	if debug == "*" {
		return true
	}

	debugList := strings.Split(debug, ",")

	for _, pattern := range debugList {
		if strings.HasPrefix(pattern, "-") {
			exclusionPattern := pattern[1:]
			if matchPattern(l.namespace, exclusionPattern) {
				return false
			}
		} else if matchPattern(l.namespace, pattern) {
			return true
		}
	}

	return false
}

func matchPattern(namespace, pattern string) bool {
	if strings.Contains(pattern, "*") {
		pattern = strings.Replace(pattern, "*", "", -1)
		return strings.HasPrefix(namespace, pattern)
	}

	return namespace == pattern
}

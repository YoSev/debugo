package debugo

import (
	"regexp"
	"strings"
)

func (l *Debugger) matchNamespace() bool {
	if l.forced || debug == "*" {
		return true
	}

	debugList := strings.Split(debug, ",")

	// Separate the exclusion and inclusion patterns
	var exclusionPatterns []string
	var inclusionPatterns []string

	for _, pattern := range debugList {
		pattern = strings.ToLower(strings.TrimSpace(pattern))
		if strings.HasPrefix(pattern, "-") {
			exclusionPatterns = append(exclusionPatterns, pattern[1:]) // Remove the "-" and store it as an exclusion
		} else {
			inclusionPatterns = append(inclusionPatterns, pattern)
		}
	}

	// Check if any exclusion pattern matches the namespace
	for _, exclusionPattern := range exclusionPatterns {
		if matchPattern(l.namespace, exclusionPattern) {
			return false // If an exclusion matches, return false immediately
		}
	}

	// Check if any inclusion pattern matches the namespace
	for _, inclusionPattern := range inclusionPatterns {
		if matchPattern(l.namespace, inclusionPattern) {
			return true // If an inclusion matches, return true
		}
	}

	return false
}

func matchPattern(namespace, pattern string) bool {
	// Replace '*' with '.*' for regex matching (.* matches any sequence of characters)
	regexPattern := "^" + strings.ReplaceAll(pattern, "*", ".*") + "$"

	// Compile the pattern into a regular expression
	re, err := regexp.Compile(regexPattern)
	if err != nil {
		return false
	}

	// Check if the namespace matches the regular expression
	return re.MatchString(namespace)
}

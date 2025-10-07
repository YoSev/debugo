package debugo

import (
	"regexp"
	"strings"
)

func (d *Debugger) matchNamespace() bool {
	namespace := GetNamespace()
	if namespace == "*" {
		return true
	}

	debugList := strings.Split(namespace, ",")

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
		if matchPattern(d.namespace, exclusionPattern) {
			return false // If an exclusion matches, return false immediately
		}
	}

	// Check if any inclusion pattern matches the namespace
	for _, inclusionPattern := range inclusionPatterns {
		if matchPattern(d.namespace, inclusionPattern) {
			return true // If an inclusion matches, return true
		}
	}

	return false
}

func matchPattern(namespace, pattern string) bool {
	// Handle the "optional" case with ":?"
	if strings.HasSuffix(pattern, ":?") {
		base := strings.TrimSuffix(pattern, ":?")
		// Match exactly the base or the base followed by anything
		regexPattern := "^" + regexp.QuoteMeta(base) + "(:.*)?$"
		re, _ := regexp.Compile(regexPattern) // can not fail due to QuoteMeta
		return re.MatchString(namespace)
	}

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

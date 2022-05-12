package logs

import "strings"

var applicationMap = map[rune]string{
	'❗': "recommendation",
	'🔍': "search",
	'☀': "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	idx := -1
	app := "default"
	for k, v := range applicationMap {
		if i := strings.IndexRune(log, k); (i < idx || idx == -1) && i >= 0 {
			idx = i
			app = v
		}
	}
	return app
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string([]rune{oldRune}), string([]rune{newRune}))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}

package parsinglogfiles

import (
	"regexp"
	"strings"
)

var (
	logLineRegex  = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`)
	sepRegex      = regexp.MustCompile(`<[-~=*]*>`)
	passwordRegex = regexp.MustCompile(`(?i)".*password.*"`)
	eolRegex      = regexp.MustCompile(`end-of-line\d*`)
	userLogRegex  = regexp.MustCompile(`User\s+[a-zA-Z0-9]+`)
)

func IsValidLine(text string) bool {
	return logLineRegex.MatchString(text)
}

func SplitLogLine(text string) []string {
	if text == "" {
		return []string{""}
	}
	return sepRegex.Split(text, len(text))
}

func CountQuotedPasswords(lines []string) int {
	result := 0
	for _, l := range lines {
		if passwordRegex.MatchString(l) {
			result++
		}
	}
	return result
}

func RemoveEndOfLineText(text string) string {
	return eolRegex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	res := make([]string, len(lines))
	for i := range lines {
		if s := userLogRegex.FindString(lines[i]); s != "" {
			res[i] = "[USR] " + strings.TrimSpace(strings.ReplaceAll(s, "User", "")) + " " + lines[i]
			continue
		}
		res[i] = lines[i]
	}
	return res
}

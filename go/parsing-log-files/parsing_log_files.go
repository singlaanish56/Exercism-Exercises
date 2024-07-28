package parsinglogfiles

import (
	"strings"
	"regexp"
	"fmt"
)
func IsValidLine(text string) bool {
	validPrefixes := []string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}
	for _, prefix := range validPrefixes {
		if strings.HasPrefix(text, prefix) {
			return true
		}
	}
	return false
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-~*=]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`"([^"]*password[^"]*)"`)

	for _, line := range lines {
		if re.MatchString(strings.ToLower(line)) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	taggedLines := make([]string, len(lines))

	for i, line := range lines {
		if matches := re.FindStringSubmatch(line); matches != nil {
			user := matches[1]
			taggedLines[i] = fmt.Sprintf("[USR] %s %s", user, line)
		} else {
			taggedLines[i] = line
		}
	}

	return taggedLines
}

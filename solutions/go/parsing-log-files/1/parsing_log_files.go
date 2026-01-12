package parsinglogfiles
import "regexp"

func IsValidLine(text string) bool {
	return regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`).MatchString(text)
}

func SplitLogLine(text string) []string {
    return regexp.MustCompile(`<[~*=-]*>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
	c := 0
	for _, l := range lines {
		if re.MatchString(l) {
			c++
		}
	}
	return c
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d+`).ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    re := regexp.MustCompile(`User\s+(\S+)`)
	for i, l := range lines {
		if m := re.FindStringSubmatch(l); m != nil {
			lines[i] = "[USR] " + m[1] + " " + l
		}
	}
	return lines
}

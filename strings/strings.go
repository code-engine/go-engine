package strings

import (
	"bufio"
	"regexp"
	"strings"
)

func MultiLineToSlice(str string) []string {
	scanner := bufio.NewScanner(strings.NewReader(str))

	out := []string{}

	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return out
}

func Extract(rgx, str string) ([]string, error) {
	re, err := regexp.Compile(rgx)

	if err != nil {
		return []string{}, err
	}

	submatches := re.FindStringSubmatch(str)

	if len(submatches) <= 0 {
		return []string{}, nil
	}

	return submatches[1:], nil
}

package vowels

import "regexp"

func IsAlt(input string) bool {
	nonAltPattern := "[aeiou]{2}|[^aeiou]{2}"
	match, _ := regexp.MatchString(nonAltPattern, input)
	return !match
}

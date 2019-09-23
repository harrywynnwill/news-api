package api

import "regexp"

// for now both the query parmas are alpha with no spaces
func IsValidQueryParams(params ...string) bool {
	for _, p := range params {
		isAlpha := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(p)
		if !isAlpha {
			return false
		}
	}
	return true
}

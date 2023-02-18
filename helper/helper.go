package helper

import "regexp"

func IsPhone(phone string) bool {
	re := regexp.MustCompile(`^0\d{2}-\d{3}-\d{4}$`)
	if !re.MatchString(phone) {
		return false
	}
	// if phone == "" {
	// 	return false
	// }
	return true
}

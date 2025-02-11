package email

import (
	"net/mail"
	"regexp"
)

func isValidEmailFormat(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

func isValidEmailByNet(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func IsValidEmail(email string) bool {
	if isValidEmailFormat(email) && isValidEmailByNet(email) {
		return true
	}
	return false
}

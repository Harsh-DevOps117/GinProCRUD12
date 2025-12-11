package validator

import (
	"fmt"
	"regexp"
)

func PasswordValidator(password string) bool {
	if len(password) < 8 {
		fmt.Println("Password must be at least 8 characters long")
		return false
	}
	hasLetter := regexp.MustCompile(`[A-Za-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	hasValidChars := regexp.MustCompile(`^[A-Za-z\d]+$`).MatchString(password)
	return hasLetter && hasDigit && hasValidChars
}

package validator

import "regexp"

func EmailValidator(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`)
	return regex.MatchString(email)
}

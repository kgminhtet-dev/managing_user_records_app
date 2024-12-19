package handler

import (
	"github.com/google/uuid"
	"regexp"
)

func isUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

func isPassword(password string) bool {
	return len(password) >= 8
}

func isEmail(email string) bool {
	var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

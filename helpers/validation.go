package helpers

import (
	"regexp"
	"strings"
	"unicode"
)

func IsValidEmail(email string) bool {
	email = strings.TrimSpace(email)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// IsStrongPassword validates password strength
// Requirements: min 8 chars, at least 1 uppercase, 1 lowercase, 1 number, 1 special char
func IsStrongPassword(password string) (bool, string) {
	if len(password) < 8 {
		return false, "password must be at least 8 characters"
	}
	if len(password) > 128 {
		return false, "password must not exceed 128 characters"
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return false, "password must contain at least one uppercase letter"
	}
	if !hasLower {
		return false, "password must contain at least one lowercase letter"
	}
	if !hasNumber {
		return false, "password must contain at least one number"
	}
	if !hasSpecial {
		return false, "password must contain at least one special character"
	}

	return true, ""
}

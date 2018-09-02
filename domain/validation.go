package domain

import "regexp"

// ValidateEmail common function
func ValidateEmail(email string) bool {
	r := regexp.MustCompile(`^[a-z0-9._+\-]+@[a-z0-9.\-]+\.[a-z]{2,8}$`)
	return !r.MatchString(email)
}

// ValidateURL common function
func ValidateURL(avatarURL string) bool {
	r := regexp.MustCompile(`^https?://[a-z0-9.\-\/]+$`)
	return !r.MatchString(avatarURL)
}

package main
const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	special = "!@#$%^&*"
)

func GeneratePassword(length int, useUppercase, useDigits, useSpecial bool) string {
	var alphabet = lowercase
	if useUppercase {
		alphabet += uppercase
	}
	if useDigits {
		alphabet += digits
	}
	if useSpecial {
		alphabet += special
	}
	len := len(alphabet)
	var result string
	for i := 0; i < length; i++ {
		ind := i % len
		result += string(alphabet[ind])
	}
	return result
}

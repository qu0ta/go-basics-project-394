package main
const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	special = "!@#$%^&*"
)
func NextRandom(number int) int {
    return (16807 * number) % 2147483647
}
func GeneratePassword(length, seed int, useUppercase, useDigits, useSpecial bool) string {
	if length <= 0 {
		return ""
	}
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
	curr := seed
	len := len(alphabet)
	var result string
	for range len {

		ind := curr % len
		result += string(alphabet[ind])
		curr = NextRandom(curr)
	}
	return result
}

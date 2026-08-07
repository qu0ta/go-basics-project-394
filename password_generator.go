package main

const lowercase = "abcdefghijklmnopqrstuvwxyz"
const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits    = "0123456789"
func GeneratePassword(length int, useUppercase, useDigits bool) string {
	var alphabet = lowercase
	if useUppercase {
		alphabet += uppercase
	}
	if useDigits {
		alphabet += digits
	}
	len := len(alphabet)
	var result string
	for i := 0; i < length; i++ {
		ind := i % len
		result += string(alphabet[ind])
	}
	return result
}

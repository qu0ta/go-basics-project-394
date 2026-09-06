package main

import (
	"fmt"
	"strings"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	special   = "!@#$%^&*"
)

var verdicts = map[int]string{
	0: "Слабый",
	1: "Слабый",
	2: "Слабый",
	3: "Средний",
	4: "Надёжный",
	5: "Очень надёжный",
}

// NextRandom возвращает следующее псевдослучайное число.
func NextRandom(number int) int {
	return (16807 * number) % 2147483647
}

// GeneratePassword генерирует пароль заданной длины на основе seed.
func GeneratePassword(length, seed int, useUppercase, useDigits, useSpecial bool) string {
	alphabet := lowercase

	if useUppercase {
		alphabet += uppercase
	}
	if useDigits {
		alphabet += digits
	}
	if useSpecial {
		alphabet += special
	}

	current := NextRandom(seed)
	alphabetLength := len(alphabet)

	var password strings.Builder

	for i := 0; i < length; i++ {
		index := current % alphabetLength
		password.WriteByte(alphabet[index])
		current = NextRandom(current)
	}

	return password.String()
}

// CheckPassword возвращает оценку надёжности пароля.
func CheckPassword(password string) string {
	score := strengthScore(password)

	return fmt.Sprintf(
		"%s пароль (оценка %d из 5)",
		verdicts[score],
		score,
	)
}

func strengthScore(password string) int {
	score := 0

	if len(password) >= 8 {
		score++
	}
	if has(password, lowercase) {
		score++
	}
	if has(password, uppercase) {
		score++
	}
	if has(password, digits) {
		score++
	}
	if has(password, special) {
		score++
	}

	return score
}

func has(password, chars string) bool {
	for _, char := range password {
		if strings.ContainsRune(chars, char) {
			return true
		}
	}

	return false
}

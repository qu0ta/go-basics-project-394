package main

import "fmt"

const (
    lowercase = "abcdefghijklmnopqrstuvwxyz"
    uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    digits    = "0123456789"
    special   = "!@#$%^&*"
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
    alphabetLen := len(alphabet)
    var result string
    
    for i := 0; i < length; i++ {
        ind := curr % alphabetLen
        result += string(alphabet[ind])
        curr = NextRandom(curr)
    }
    return result
}


func CheckPassword(password string) string {
	var score int
	if len(password) >= 8 {
		score++
	}
	for _, l := range lowercase {
		isLowercased := false
		for _, p := range password {
			if l == p {
				score++
				isLowercased = true
				break
			}
		}
		if isLowercased {
			break
		}
	}
	for _, l := range uppercase {
		isUppercased := false
		for _, p := range password {
			if l == p {
				score++
				isUppercased = true
				break
			}
		}
		if isUppercased {
			break
		}
	}
	for _, l := range digits {
		isdigitted := false
		for _, p := range password {
			if l == p {
				score++
				isdigitted = true
				break
			}
		}
		if isdigitted {
			break
		}
	}
	for _, l := range special {
		isSpecialized := false
		for _, p := range password {
			if l == p {
				score++
				isSpecialized = true
				break
			}
		}
		if isSpecialized {
			break
		}
	}
	m := make(map[int]string)
	m[0] = "Слабый"
	m[1] = "Слабый"
	m[2] = "Слабый"
	m[3] = "Средний"
	m[4] = "Надёжный"
	m[5] = "Очень надёжный"

	return fmt.Sprintf("%s пароль (оценка %d из 5)", m[score], score)

}
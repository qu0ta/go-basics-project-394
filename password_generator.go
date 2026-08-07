package main

const lowercase = "abcdefghijklmnopqrstuvwxyz"

func GeneratePassword(length int) string {
	len := len(lowercase)
	var result string
	for i := 0; i < length; i++ {
		ind := i % len
		result += string(lowercase[ind])
	}
	return result
}

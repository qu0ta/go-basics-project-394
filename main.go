package main

import (
	"fmt"
)

func main() {
	fmt.Println(GeneratePassword(-3, 42, true, true, false))
}

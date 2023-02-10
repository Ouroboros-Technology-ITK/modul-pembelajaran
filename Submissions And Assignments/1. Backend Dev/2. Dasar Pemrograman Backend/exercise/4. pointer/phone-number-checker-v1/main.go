package main

import (
	"fmt"
)

func PhoneNumberChecker(number string, result *string) {
	// TODO: answer here
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = ""
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}

package main

import (
	"fmt"
)

func Readfile(path string) ([]string, error) {
	return nil, nil // TODO: replace this
}

func CalculateProfitLoss(data []string) string {
	return "" // TODO: replace this
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}

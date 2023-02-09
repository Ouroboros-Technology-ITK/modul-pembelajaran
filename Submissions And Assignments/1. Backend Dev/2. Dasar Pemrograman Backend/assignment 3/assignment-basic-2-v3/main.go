package main

import (
	"fmt"
	"strconv"
)

// Optional, kalian bisa membuat fungsi helper seperti berikut, untuk menerapkan DRY principle
// fungsi ini akan mengubah int ke currency Rupiah
// example: int 1000 => Rp 1.000
func FormatRupiah(number int) string {
	var result string

	numStr := strconv.Itoa(number)

	for i := len(numStr) - 1; i >= 0; i-- {
		result = string(numStr[i]) + result
		if (len(numStr)-i)%3 == 0 && i != 0 {
			result = "." + result
		}
	}

	return "Rp " + result
}

func GetDayDifference(date string) int {
	return 0 // TODO: replace this
}

func GetSalary(rangeDay int, data [][]string) map[string]string {
	return nil // TODO: replace this
}

func GetSalaryOverview(dateRange string, data [][]string) map[string]string {
	return nil // TODO: replace this
}

func main() {
	res := GetSalaryOverview("21 February - 23 February 2021", [][]string{
		{"andi", "Rojaki", "raji", "supri"},
		{"andi", "Rojaki", "raji"},
		{"andi", "raji", "supri"},
		{"andi", "Rojaki", "raji", "supri"},
	})

	fmt.Println(res)
}

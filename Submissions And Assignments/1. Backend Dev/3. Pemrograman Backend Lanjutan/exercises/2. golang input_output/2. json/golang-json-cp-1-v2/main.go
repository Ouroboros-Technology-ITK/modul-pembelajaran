package main

import (
	"fmt"
)

type Report struct {
	// TODO: answer here
}


// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
	// TODO: answer here
}

func GradePoint(report Report) float64 {
	return 0.0 // TODO: replace this
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}

package main

import "fmt"

type School struct {
}

func (s *School) AddGrade(grades ...int) {
}

func Analysis(s School) (float64, int, int) {
	nil, nil, nil // TODO: replace this
}

// gunakan untuk melakukan debugging
func main() {
	avg, min, max := Analysis(School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{100, 90, 80, 70, 60},
	})

	fmt.Println(avg, min, max)
}

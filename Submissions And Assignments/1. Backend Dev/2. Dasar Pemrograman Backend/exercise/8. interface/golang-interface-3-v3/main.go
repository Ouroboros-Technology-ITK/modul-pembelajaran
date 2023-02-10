package main

import (
	"fmt"
)

type Time struct {
	Hour   int
	Minute int
}


func ChangeToStandartTime(time interface{}) string {
	return "" // TODO: replace this
}

func main() {
	fmt.Println(ChangeToStandartTime("16:00"))
	fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}

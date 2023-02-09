package main

type Employee interface {
	GetBonus() float64
}

type Junior struct {
}

type Senior struct {
}

type Manager struct {
}


func EmployeeBonus(employee Employee) float64 {
	return 0.0 // TODO: replace this
}

func TotalEmployeeBonus(employees []Employee) float64 {
	return nil // TODO: replace this
}

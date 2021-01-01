package main

import "fmt"

type workingHours int

type FullTimeEmp struct {
	empId int
	name  string
	workingHours
	salary // anonymous fields
}

type ContractEmp struct {
	empId    int
	name     string
	Duration int
	salary
}

type salary struct {
	pf      int
	basePay int
}

func main() {

	ft := FullTimeEmp{
		empId:        101,
		name:         "Raja",
		workingHours: 8,
		salary: salary{
			pf:      2000,
			basePay: 40000,
		},
	}

	ct := ContractEmp{
		empId:    102,
		name:     "John",
		Duration: 3,
		salary: salary{
			pf:      1000,
			basePay: 2000,
		},
	}

	fmt.Println(ft)
	fmt.Println(ct)

	fmt.Println(ct.basePay)

}

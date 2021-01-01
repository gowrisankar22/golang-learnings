package main

import "fmt"

type FullTimeEmp struct {
	empId        int
	name         string
	workingHours int
	sal          salary
}

type ContractEmp struct {
	empId    int
	name     string
	Duration int
	sal      salary
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
		sal: salary{
			pf:      2000,
			basePay: 40000,
		},
	}

	ct := ContractEmp{
		empId:    102,
		name:     "John",
		Duration: 3,
		sal: salary{
			pf:      1000,
			basePay: 2000,
		},
	}

	fmt.Println(ft)
	fmt.Println(ct)

	fmt.Println(ct.sal.basePay)

}

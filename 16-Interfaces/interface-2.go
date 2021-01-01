package main

import "fmt"

type Expense interface {
	CalculateSalary() int
}

type perm struct {
	empId    int
	basicPay int
	pf       int
}

type contract struct {
	empId    int
	basicPay int
	bonus    int
}

func (p perm) CalculateSalary() int {
	return p.basicPay + p.pf
}

func (c contract) CalculateSalary() int {
	return c.basicPay + c.bonus
}

func totalExpense(e ...Expense) {
	total := 0

	for _, v := range e {
		total = total + v.CalculateSalary()
	}

	fmt.Println("Total Expense ", total)

}

func main() {

	e1 := perm{
		empId:    101,
		basicPay: 20000,
		pf:       10000,
	}

	e2 := perm{
		empId:    102,
		basicPay: 25000,
		pf:       15000,
	}

	e3 := contract{
		empId:    103,
		basicPay: 22000,
		bonus:    18000,
	}

	totalExpense(e1, e2, e3)

}

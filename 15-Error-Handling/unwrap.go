package main

import (
	"errors"
	"fmt"
)

var FeesNotSubmitted = errors.New("fees Not Submitted")

var AdmissionCancelled = errors.New("admission not possible")

var foo1 = errors.New("foo")

func fees() error {
	return fmt.Errorf("%w", FeesNotSubmitted)
}
func admission() error {
	return fmt.Errorf("%w %v", fees(), AdmissionCancelled)
}

func foo() error {
	return fmt.Errorf("%w %v", admission(), foo1)
}

func main() {
	err := foo()

	if errors.Is(err, FeesNotSubmitted) {
		fmt.Println("present")
	}else {
		fmt.Println("not")
	}

	fmt.Println(err)

	err = errors.Unwrap(err)
	fmt.Println(err)

	err = errors.Unwrap(err)
	fmt.Println(err)
}

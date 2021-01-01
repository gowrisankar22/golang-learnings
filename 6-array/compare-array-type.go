package main

type name [10]string
type age [10]string

func main() {

	//var class1 name
	class := name{"abc", "xyz"}
	classAge := age{}

	_ = class
	_ = classAge

	//if name == age{
	//
	//}

	// Not possible as both are different types
}

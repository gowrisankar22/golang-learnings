package main

import "fmt"

func main() {

	class := map[int][]string{
		10: {"xyz", "abc"},
		11: {"tyu", "iou", "poi"},
	}

	fmt.Println(class[10])
	for k, v := range class {
		fmt.Println(k, v)
	}

}

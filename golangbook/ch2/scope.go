package ch2

import "fmt"

// ShowScope is
func ShowScope() {
	x := "hello"          // fuction scope
	for _, x := range x { // implicit for scope
		x := x + 'A' - 'a' // explicit for block scope
		fmt.Printf("%c", x)
	}
}

// ScopeInIf is
func ScopeInIf() {
	if s := anyFn(); s == 0 {
		fmt.Println(s)
	} else if y := anyOtherFn(s); s == y {
		fmt.Println(s, y)
	} else {
		fmt.Println(s, y)
	}

	fmt.Println(s, y) // Compilation error. x and y are not enabled here...
}

func anyFn() int {
	return 0
}

func anyOtherFn(s int) int {
	return s
}

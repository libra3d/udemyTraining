package main

import "fmt"

// switch on types
// -- normally we switch on value of variable
// -- go allows you to switch on type of variable

type Contact struct {
	greeting string
	name     string
}

// SwitchOnType works with interfaces
// we'll learn more about interfaces later
func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unkown")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("Libra")
	var t = Contact{"Good to see you,", "Tim"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
}

package main

import "fmt"

//create a new interface with type Test with a method Tester()
type Test interface {
	Tester()
}

//create a float64 Type
type MyFloat float64

//implement Tester method with type MyFloat
func (m MyFloat) Tester() {
	fmt.Println(m)
}

//print out the interface type and value
func describe(t Test) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func main() {
	var t Test         //create a new interface, that is not initiliased
	f := MyFloat(89.7) //create and assigne a new float64 variable type
	t = f              // assign f value to the interface, now t can call method Teste
	describe(t)
	t.Tester()
}

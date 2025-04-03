package main

import "fmt"

func main() {
	age := 26 //regular variable

	var agePointer *int

	agePointer = &age

	fmt.Println("Age Pointer:", *agePointer)

	fmt.Println("Age:", age)

	getAdultYears(agePointer)
	fmt.Println("Directly mutated value", age)

}

func getAdultYears(age *int) {
	// return *age - 18
	*age -= 18
}

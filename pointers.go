package main

import "fmt"

func main() {
	age := 26 //regular variable

	var agePointer *int

	agePointer = &age

	fmt.Println("Age Pointer:", *agePointer)

	fmt.Println("Age:", age)

	adultYears := getAdultYears(age)
	fmt.Println(adultYears)

}

func getAdultYears(age int) int {
	return age - 18
}

package main

import (
	"fmt"
)

func main() {
	/*
		fmt.Print("Enter your age: ")
		var age int
		fmt.Scanf("%d", &age)
	*/

	//if boolean_expression {
	//logic or statement here
	//}
	/*
		if age < 3 { //0 to 2

			fmt.Println("infant")

		} else if age > 2 && age < 13 { //2 to 12

			fmt.Println("children")

		} else if age > 12 && age <= 19 {

			fmt.Println("teen age")
		} else {
			fmt.Println("adult")
		}
	*/
	//fmt.Println(age)

	//fixed value
	/*
		switch age {
		case 2:
			fmt.Println("infant")
			fallthrough

		case 3, 4, 5, 6, 7, 8, 9, 110, 11, 12:
			fmt.Println("children")

		case 13, 14, 15, 16, 17, 18, 19:
			fmt.Println("teen age")

		default:
			fmt.Println("adult")

		}
	*/

	//for loop
	/*
		for i := 1; i <= 9; i++ {
			fmt.Println(i)
		}
	*/
	//array string literals
	/*
		students := []string{"Noyon", "karim", "Rahim"}

		for i, std := range students {

			fmt.Println(i, std)
		}
	*/

	i := 0
	for {
		fmt.Println(i, "Hello")
		i++
	}
}

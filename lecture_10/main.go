package main

import (
	"fmt"
	//"reflect"
)

func main() {

	//var students [3]string
	//students[0] = "karim"
	//students[1] = "rahim"
	//students[2] = "sofi"

	//arry
	//fmt.Println(students[2])

	//students := [...]string{"Karim", "rahim", "Sofi", "rafi", "rakib"}
	//fmt.Println(students, len(students))

	//slice

	//x := students[0:3]

	//x := make([]string, 0)

	//var fruits []string
	//fruits = append(fruits, "apple", "banana")
	//fmt.Println(fruits, len(fruits))

	//fmt.Printf("%T\n", fruits)
	//fmt.Printf("%T", students)

	//b := reflect.TypeOf(fruits).Kind().String()
	//fmt.Println(b)

	//Map

	//var x map[string]string

	x := make(map[string]string)
	x["name"] = "Noyan"
	x["height"] = "5.8"
	x["address"] = "Habiganj"

	fmt.Println(x["name"])

}

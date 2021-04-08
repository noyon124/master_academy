package main

import (
	"fmt"
)

func main() {

	//array
	//var students [3]string
	//students[0] = "asgor"
	//students[1] = "mainul"
	//students[2] = "anni"

	//fmt.Println(students)
	//fmt.Println(len(students))
	//fmt.Println(students[0])

	//sort hand way, string literals
	//students := [...]string{"asgor", "mainul", "anni", "rakib"}
	//fmt.Println(students, len(students))

	//slice

	//var students [3]string
	//students[0] = "asgor"
	//students[1] = "mainul"
	//students[2] = "anni"

	//x := make([]string, 0)
	//var furits []string
	//furits = append(furits, "apple", "banana", "mango")

	//fmt.Println(furits, len(furits))
	//fmt.Printf("%T\n", furits)
	//fmt.Printf("%T", students)

	//b := reflect.TypeOf(furits).Key().Kind().String()
	//fmt.Println(b)

	//map
	x := make(map[string]string)
	x["name"] = "noyon"

	fmt.Println(x)
}

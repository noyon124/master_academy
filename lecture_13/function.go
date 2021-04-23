package main

import "fmt"

//example-1

/*
func add(x int, y int) int {

	r := x + y

	return r
}
*/
//example-2
/*
func add(x, y int) int {

	r := x + y
	return r

}
*/

//example-3
/*
func add(x, y int) (r int) {

	r = x + y
	return r

}
*/

//exapmle-4
/*
func add(x, y int) (r int) {

	r = x + y
	return

}
*/
/*
func rectangle(l int, b int) (area int, perametar int) {
	perametar = 2 * (1 + b)
	area = 1 * b
	return
}
*/

//pointer
/*
func update(a *int, t *string) {
	*a = *a + 5
	*t = *t + "Doe"
	return
}
*/

func main() {

	//x := add(10, 30)

	//a, p := rectangle(10, 30)
	//fmt.Println(a, p)

	//pointer:

	//number := 10
	//name := "Noyon"
	//update(&number, &name)
	//fmt.Println(number, name)

	a := func(x, y int) (r int) {
		r = x * y
		return
	}(10, 10)

	//fmt.Println(a(10, 10))
	fmt.Println(a)

}

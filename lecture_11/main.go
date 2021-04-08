package main

import "fmt"

var w1 Weight

/*type Book struct {
	Title  string
	Author string
	ISBN   string
	Price  float32
	Pages  int
}*/

func main() {

	//var b1 Book
	//b1.Title = "An introduction to Programmi in go"
	//b1.Author = "CALEB DOXY"
	//b1.ISBN = "925-4658948256"
	//b1.Price = 10 - 50
	//b1.Pages = 165

	//fmt.Println(b1)
	//fmt.Println(b1.Title)
	//fmt.Println(b1.Author)
	//fmt.Println(b1.ISBN)
	//fmt.Println(b1.Price)
	//fmt.Println(b1.Pages)

	//Struct literals in Annoonymous way

	/*
		b1 := struct {
		    Titel  string
			Author string
			ISBN   string
			Price  float32
			Pages  intrle
		}{
			Titel:  "An introduction to Programming in go",
			Author: "Caleb Doxy",
			ISBN:   "925-4658948256",
			Price:  10.50,
			Pages:  165,
		}*/

	//fmt.Println(b1)

	var w1 Weight
	w1 = 30.50
	fmt.Println(w1, name)

}

package main

import (
	"fmt"

	. "./mypag"
)

func main() {
	fmt.Printf("%s\n", "hello")
	//var a string = "1"
	//var b string = "1"
	//var p1 Person = Person{Name: "p1", Age: 1}
	var p1 Person = Person{Name: "p1", Age: 1}
	var p2 Person = Person{Name: "p2", Age: 2}
	//
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(Baba(&p1, &p2))
}

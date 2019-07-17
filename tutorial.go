package main

import "fmt"

func main() {
	// fmt.Println("hello world")
	// fmt.Println("1+1 =", 1+1)
	// fmt.Println(true && false)
	// var a = "initial"
	// f:= "apple"
	// fmt.Println(a, f)

	// var (
	// 	name     = "Prince Oberyn"
	// 	age      = 32
	// 	location = "Dorne"
	// )
	// fmt.Printf("%s (%d) of %s \n", name, age, location)

	name, location := "Prince Oberyn", "Dorne"
	age := 32
	fmt.Printf("%s (%d) of %s \n", name, age, location)

	action := func() {
		fmt.Println("Action function called")
	}
	action()

	const Pi = 3.14
	const (
		StatusOK                   = 200
		StatusCreated              = 201
		StatusAccepted             = 202
		StatusNonAuthoritativeInfo = 203
		StatusNoContent            = 204
		StatusResetContent         = 205
		StatusPartialContent       = 206
	)
	fmt.Println("Constant Pi equals:", Pi)
}

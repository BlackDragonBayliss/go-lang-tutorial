package main

import "fmt"
import "math"

const s string = "constant"

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

	// name, location := "Prince Oberyn", "Dorne"
	// age := 32
	// fmt.Printf("%s (%d) of %s \n", name, age, location)

	// action := func() {
	// 	fmt.Println("Action function called")
	// }
	// action()

	// const Pi = 3.14
	// const (
	// 	StatusOK                   = 200
	// 	StatusCreated              = 201
	// 	StatusAccepted             = 202
	// 	StatusNonAuthoritativeInfo = 203
	// 	StatusNoContent            = 204
	// 	StatusResetContent         = 205
	// 	StatusPartialContent       = 206
	// )
	// fmt.Println("Constant Pi equals:", Pi)

	fmt.Println(s)

	// A `const` statement can appear anywhere a `var`
	// statement can.
	const n = 500000000

	// Constant expressions perform arithmetic with
	// arbitrary precision.
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it's given
	// one, such as by an explicit conversion.
	fmt.Println(int64(d))

	// A number can be given a type by using it in a
	// context that requires one, such as a variable
	// assignment or function call. For example, here
	// `math.Sin` expects a `float64`.
	fmt.Println(math.Sin(n))

}

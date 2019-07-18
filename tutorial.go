package main

import "fmt"

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

	// fmt.Println(s)

	// // A `const` statement can appear anywhere a `var`
	// // statement can.
	// const n = 500000000

	// // Constant expressions perform arithmetic with
	// // arbitrary precision.
	// const d = 3e20 / n
	// fmt.Println(d)

	// // A numeric constant has no type until it's given
	// // one, such as by an explicit conversion.
	// fmt.Println(int64(d))

	// // A number can be given a type by using it in a
	// // context that requires one, such as a variable
	// // assignment or function call. For example, here
	// // `math.Sin` expects a `float64`.
	// fmt.Println(math.Sin(n))

	// The most basic type, with a single condition.
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// j := 0
	// for j <= 5 {
	// 	fmt.Println(j)
	// 	j = j + 1
	// }

	// for n := 0; n <= 5; n++ {
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }

	// for n := 0; n <= 5; n++ {
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }

	// for {
	// 	fmt.Println("loop")
	// 	break
	// }

	if num := -10; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	// // A classic initial/condition/after `for` loop.
	// for j := 7; j <= 9; j++ {
	// 	fmt.Println(j)
	// }

	// // `for` without a condition will loop repeatedly
	// // until you `break` out of the loop or `return` from
	// // the enclosing function.
	// for {
	// 	fmt.Println("loop")
	// 	break
	// }

	// // You can also `continue` to the next iteration of
	// // the loop.
	// for n := 0; n <= 5; n++ {
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }
}

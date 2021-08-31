package main

import "fmt"

// A normal function, this cannot be nested
func standardFunction(arg1 string, arg2 string) string {
	return fmt.Sprintf("%s:%s", arg1, arg2)
}

func main() {

	// A function assigned to a variable
	var assignedFunction1 = standardFunction

	// An anonymous function assigned to a variable and nested
	var assignedFunction2 = func(arg1 string, arg2 string) string {
		return fmt.Sprintf("%s:%s", arg1, arg2)
	}

	// A higher-order-function that accepts functions as argument and returns a function
	var functionAsArgumentAndReturn = func(addFn func(int, int) int, arg1 int, arg2 int) func(int) int {
		var out = addFn(arg1, arg2)
		// This returns a closure
		return func(numArg int) int {
			return out + numArg
		}
	}

	var out = functionAsArgumentAndReturn(
		func(a, b int) int {
			return a + b
		},
		5,
		10,
	)(10)
	fmt.Println(out) // prints 25

	// Nested anonymous functions
	var nested = func() {
		fmt.Println("outer fn")
		var nested2 = func() {
			fmt.Println("inner fn")
			var nested3 = func() {
				fmt.Println("inner arrow")
			}
			nested3()
		}
		nested2()
	}

	nested() // prints:
	// outer fn
	// inner fn
	// inner arrow

	// this is a higher-order-function that returns a function
	var add = func(x int) func(y int) int {
		// A function is returned here as closure
		// variable x is obtained from the outer scope of this method and memorized in the closure
		return func(y int) int {
			return x + y
		}
	}

	// we are currying the add method to create more variations
	var add10 = add(10)
	var add20 = add(20)
	var add30 = add(30)

	fmt.Println(add10(5)) // 15
	fmt.Println(add20(5)) // 25
	fmt.Println(add30(5)) // 35

	// An anonymous function invoked immediately(IIFE)
	(func() {
		fmt.Println("anonymous fn")
	})()
	// prints: anonymous fn

	assignedFunction1("a", "b")
	assignedFunction2("a", "b")
}
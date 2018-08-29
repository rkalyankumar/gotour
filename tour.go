package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

func add(a, b int) int {
	return a + b
}

func addandsub(a, b int) (int, int) {
	return add(a, b), a - b
}

/**
 * A Naked return function Example:
 * 	A Naked return function is one that contains named return values.
 *  These variables are treated as definitions in the top of the function.
 *  The function implementing a naked return simply does a "return" at the end
 *  to return the named variables
 */

func split(sum int) (x, y int) {
	x = sum * 3 / 9
	y = sum - x
	return
}

// Some variables
var i, j int = 1, 2

// Block style var declarations

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// const can't be declared using := & they can be character, string, boolean or numeric values
const MyPi = 3.14

func sum() int {
	s := 0
	for i := 0; i < 10; i++ {
		s += i
	}
	return s
}

func sumOptional() int {
	s := 1
	for s < 1000 {
		s += s
	}
	return s
}

func mysqrt(x float64) string {
	if x < 0 {
		return mysqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Printf("My favorite number is %d\n", rand.Intn(10))
	fmt.Printf("Sqrt(7): %g\n", math.Sqrt(7))
	fmt.Println(math.Pi)

	var c, python, java = true, false, "no!"

	fmt.Println(split(17))
	fmt.Println(i, j, c, python, java)
	k := "Kalyan" // var declaration with implicit type
	fmt.Println(k)

	fmt.Println(addandsub(2, 3))

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	fmt.Println(sum())
	fmt.Println(sumOptional())

	s := 1
	for s < 1000 {
		s += s
	}
	fmt.Println(s)
	/*
		 Indefinite looping:
		 	for {
			}
	*/
	// mysqrt() demonstrates if loop
	fmt.Println(mysqrt(2), mysqrt(-4))
	// pow() demonstrates if with a short statement
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	// switch statements
	fmt.Print("GO runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	// switch with no condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	// defer - defers the execution of a function until the surrounding function returns
	// deferred call's arguments are evaluated immediately, but the function call isn't
	// executed until the surrounding function returns
	defer fmt.Println("World!")
	fmt.Print("Hello, ")
	// stacking defers
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

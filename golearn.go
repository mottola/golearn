package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func add(x int, y int) int {
	return x + y
}

// (string, string) allows the return of two strings
func swap(x string, y string) (string, string) {
	return y, x
}

// if statement
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Printf("hello, go!\n")
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Printf("Now you have %g problems\n", math.Sqrt(7))
	fmt.Println(math.Pi)

	fmt.Println(add(42, 13))

	// := can only be used within a function and type is inferred
	a, b := swap("Go", "Hello")
	fmt.Println(a, b)

	// constants
	const World = "世界"
	const Pi = 3.14
	fmt.Println("Hello, World")
	fmt.Println("Happy", Pi, "Day")

	// looping - Go only has for for loops
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// while equivalent
	newsum := 1
	for newsum < 1000 {
		newsum += newsum
	}
	fmt.Println(newsum)

	// if statement
	fmt.Println(sqrt(2), sqrt(-4))

	// switch statment - evaluates from top to bottom, stopping when a case succeeds
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s", os)
	}

	// more case
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today!")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}

	// pointers
	name := "Rick"
	// n is a pointer to name
	n := &name
	// print name through the pointer n
	fmt.Println(*n)

	// structs - collections of fields
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{1, 2}
	// structs are accessed with dot notation
	v.X = 4
	fmt.Println(v.X)

	// arrays - fixed size in Go - size and type are declared
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "Go"
	fmt.Println(arr[0], arr[1])
	fmt.Println(arr)

	// initialized array
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slices - Go's dynamically-sized view into the elements of an array
	slc := primes[1:4]
	fmt.Println(slc)

	// modfiying an element of a slice will mutate the underlying array
	slc[0] = 10
	fmt.Println(slc)
	fmt.Println(primes)

	// a slice has both a length and a capacity
	fmt.Printf("len=%d cap=%d\n", len(slc), cap(slc))

	// the make function allocates a zeroed array and returns a slice that refers to that array
	made := make([]int, 0, 5) // slice referencing an array of ints with len 0 and cap 5
	fmt.Println(made)

	// the append function adds new elements to a slice resizing as needed
	made = append(made, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(made)

	// the range form of the for loop iterates over a slice or map
	// when ranging over a slice, two values are returned for each iteration -
	// the first is the index, and the second is a copy of the element at that index
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, e := range pow {
		fmt.Printf("Squaring value %d at index %d results in %d\n", e, i, e*e)
	}

	// for just the index, drop the value
	// for i := range pow
	// to skip the index use _
	// for _, e := range pow

	/* PICK UP AT MAPS IN THE TOUR OF GO (Page 21/27) */
}

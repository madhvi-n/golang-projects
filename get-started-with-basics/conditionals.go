package main

import (
  "fmt"
  "math"
  "time"
)

/*
The basic for loop has three components separated by semicolons:
the init statement: executed before the first iteration
the condition expression: evaluated before every iteration
the post statement: executed at the end of every iteration
*/

func forLoop() {
  sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// C's while is spelled for in Go.

func whileLoop() {
  sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// If statement

func sqrt(x float64) string {
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}

// If with short statement

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

/*
  if condition {
    // if block
  } else {
    // else block
  }
*/


func switchLoop() {
  fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func switchWithNoCondition() {
  t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	forLoop()
  whileLoop()
  switchLoop()
  fmt.Println(sqrt(2), sqrt(-4))
  fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

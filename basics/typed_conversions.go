package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

  v:= true
  fmt.Println("v is of type %T\n", v)
}

/*
The expression T(v) converts the value v to the type T.

Some numeric conversions:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
Or, put more simply:

i := 42
f := float64(i)
u := uint(f)


Type inference
When the right hand side of the declaration is typed, the new variable is of that same type:

var i int
j := i // j is an int
But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
*/

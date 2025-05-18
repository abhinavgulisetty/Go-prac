package main

import (
	"fmt"
	"reflect"
)

// For checking types

func main() {
	//Declaring variables
	var name string = "LOL"
	//Inferred type
	var age = 20
	//Short hand
	city := "New York"
	//Multiple variables
	a, b, c := 1, 2, 3
	//Multiple variables with same type
	var x, y, z float32 = 1.1, 2.0, 3.0
	//Multiple variables with different types
	var i, j, k = 1, 2.0, "3"
	//Multiple variables with different types and short hand
	var m, n, o = 1, 2.3, "3"
	//Multiple variables with different types and short hand and inferred type
	var p, q, r = 1, 2.0, "3"
	fmt.Println(name, age, city, a, b, c, x, y, z, i, j, k, m, n, o, p, q, r)
	fmt.Println(reflect.TypeOf(name), reflect.TypeOf(age), reflect.TypeOf(city), reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c), reflect.TypeOf(x), reflect.TypeOf(y), reflect.TypeOf(z), reflect.TypeOf(i), reflect.TypeOf(j), reflect.TypeOf(k), reflect.TypeOf(m), reflect.TypeOf(n), reflect.TypeOf(o), reflect.TypeOf(p), reflect.TypeOf(q), reflect.TypeOf(r))
}

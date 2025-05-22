package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["Test"] = 1
	m["Test2"] = 2
	fmt.Println(m["Test3"]) // 0 because it is not in the map
	fmt.Println(m["Test"])  // 1 because it is in the map
	fmt.Println(len(m))
	fmt.Println(m) // map[Test:1 Test2:2]
	//Delete a key from the map
	delete(m, "Test")
	fmt.Println(m) // map[Test2:2]

	//Another way to declare a map
	d := map[string]int{
		"Test":  1,
		"Test2": 2,
	}
	fmt.Println(d) // map[Test:1 Test2:2}

	v, ok := d["Test"]
	//ok is a boolean that tells if the key is in the map or not
	//If ok is true, then v is the value of the key in the map
	//If ok is false, then v is the zero value of the value type
	fmt.Println(v)
	if ok {
		fmt.Println("Key found")
	} else {
		fmt.Println("Key not found")
	}

	p1 := map[string]int{
		"Test":  1,
		"Test2": 3,
	}
	p2 := map[string]int{
		"Test":  1,
		"Test2": 3,
	}
	fmt.Println(maps.Equal(p1, p2))
	fmt.Println(maps.Equal(p1, d)) // false because they are not the same map
}

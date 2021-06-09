package main

import "fmt"

func main() {
	// maps are key / value data
	// just like json
	// maps are mutable
	mp := map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}
	// fmt.Println(mp["apple"])
	mp["apple"] = 900

	// check if key exists in a map,
	// wherein val is the value of the key
	// while ok is the logic where key is exists
	val, ok := mp["apple"]
	if ok {
		fmt.Println(val, ok)
	}
	fmt.Println(mp)
}

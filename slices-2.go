package main

import "fmt"

func main() {
	// another way to declare an array
	// myArray := [...]int{1, 2, 3}
	// fmt.Println(myArray)

	a := []int{1, 2, 3, 4, 5, 6, 4}
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }

	// for with range
	// where i = index
	// where element is the data
	// if you don't want to use the index or element
	// just replace it with the underscore
	// for _, element := range a {
	// 	fmt.Printf("%d\n", element)
	// }

	// return duplicate values in slices
	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			element2 := a[j]
			if element2 == element {
				fmt.Println(element)
			}
		}
	}
}

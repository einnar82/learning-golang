package main

import "fmt"

func main() {
	//first declaration of array
	var arr [5]int
	//access array
	arr[0] = 100
	arr[4] = 900
	fmt.Println(arr[0])

	//another declaration of array
	num := [3]int{4, 5, 6}
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}

	fmt.Println(sum)

	// 2 dimensional array
	arr2D := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr2D)
}

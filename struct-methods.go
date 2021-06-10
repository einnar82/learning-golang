package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

// function receivers
// getters and setters in structs
func (student Student) getAge() int {
	return student.age
}

// if you have a method that
// modifies a variable
// always use a pointer
func (student *Student) setAge(age int) {
	student.age = age
}

func (student Student) getMaxGrade() int {
	curMax := 0
	for _, v := range student.grades {
		if curMax < v {
			curMax = v
		}
	}
	return curMax
}

func (student Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range student.grades {
		sum += v
	}
	return float32(sum) / float32(len(student.grades))
}

func main() {
	s1 := Student{"Rannie", []int{20, 30, 40, 60}, 19}
	s2 := Student{"Rannie", []int{20, 30, 40, 60, 88}, 33}
	// fmt.Println(s1)
	// s1.setAge(28)
	// fmt.Println(s1)
	average := s1.getAverageGrade()
	average2 := s2.getMaxGrade()
	fmt.Println(average)
	fmt.Println(average2)
}

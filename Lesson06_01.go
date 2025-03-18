package main

import (
	"fmt"
)

/*
	func cal(x, y int) int {
		result := x + y
		return result
	}
*/
type Student struct {
	name          string
	math, english float64
}
type User struct {
	gender string
	age    int
}

func cal(x, y int) (r int) {
	r = x + y
	return
}
func main() {

	/*
		var s Student
		s.name = "chiba"
		s.math = 80
		s.english = 70
	*/
	var u1 User
	u1.gender = "male"
	u1.age = 20

	u2 := User{gender: "famel", age: 15}

	//s := Student{name: "chiba", math: 80, english: 70}
	/*
		array := [...]int{15, 18, 30, 50}
		for i := 0; i < len(array); i++ {
			if array[i] <= 15 {
				fmt.Println("yung")
			} else if array[i] >= 30 {
				fmt.Println("adault")
			} else {
				fmt.Println("mada yung")
			}
		}
	*/

	//fmt.Println(cal(10, 5))
	//fmt.Println(s)
	fmt.Println(u1)
	fmt.Println(u2)
}

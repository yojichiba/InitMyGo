package main

import (
	"fmt"
)

/*
	type Student struct {
		name          string
		math, english float64
	}

	func (s Student) avg() {
		fmt.Println(s.name, (s.math+s.english)/2)

}

	func main() {
		a001 := Student{name: "Sato", math: 80, english: 70}
		a001.avg()
	}
*/
/*
type Student struct {
	name string
}

func (s Student) avg(math, english float64) (nameResult string, avgResult float64) {
	nameResult = s.name
	avgResult = (math + english) / 2
	return
}

func main() {
	a001 := Student{name: "Chba"}
	fmt.Println(a001.avg(70, 90))
}
*/
/*
type User2 struct {
	name string
}

func (u2 User2) cal(weight, height float64) (result float64) {
	result = weight / height / height * 10000
	return
}
func main() {
	bmi := User2{name: "chiba"}
	fmt.Println(bmi.name, bmi.cal(77.7, 165))
}
*/
type User2 struct {
	name string
}

func (u2 User2) calAvg(data []int) (result float64) {
	sum := 0
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	result = float64(sum) / float64(len(data))
	return
}
func judge(avg float64) (result string) {
	if avg <= 45 {
		result = "missed"
	} else {
		result = "succes"
	}
	return
}
func main() {
	a001 := User2{name: "Chiba"}
	data := []int{50, 20, 35, 40, 45}
	fmt.Println(a001.name, judge(a001.calAvg(data)))
}

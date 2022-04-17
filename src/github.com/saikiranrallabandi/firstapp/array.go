package main

import (
	"fmt"
)

func main()  {
        grades := [...]string{"saikiran", "rallabandi"}
		fmt.Printf("Grades: %v\n",grades[0])
		var students [3]string
		students[0] = "aadya"
		fmt.Printf("Grades: %v",students)
		var identityMatrix [3][3]int
		identityMatrix[0] = [3]int{1,0,0}
		fmt.Println(identityMatrix)
		//a := [...]int{1,2,3}
		//b := &a
		//b[1] = 5
		//fmt.Println(a)
		//fmt.Println(b)
		//a := []int{1,2,3} //slice are ref
		a1 := []int{1,2,3,4,5,6,7,8,9}
		b := a1[:]
		c := a1[3:]
		d := a1[:6]
		e := a1[3:6]
		// a1[0] = append(a1,[]int{2,3,4}...spread operator javascript]) adding
		a1[5] = 42
		fmt.Println(a1)
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(d)
		fmt.Println(e)

		//a := make([]int, 3)
		d1 := []int{1,2,3,4,5}
		//b3 := d1[:len(d1) - 1 ]
		c3 := append(d1[:2], d1[3:]...)
		fmt.Println(c3)


}
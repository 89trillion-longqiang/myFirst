package main

import "fmt"

type add_int int

func main()  {
	/*arr := [...]int{1,2,3,4,5,6,7}
	var s1 []int
	s1 = arr[5:7]
	fmt.Println(s1)

	s2 := make([]int,3,10)
	fmt.Println(s2)
	copy(s2,s1[:2])
	fmt.Println(s2)*/
	var a add_int
	a.Increase()
	fmt.Println(a)

	s := make([]int , 0)
	s = insert(s)
	fmt.Println(s)
}

func insert(s []int) []int{
	s = append(s,1,2,3)
	return s
}
func (a *add_int)Increase(){
	*a = *a + 100
}
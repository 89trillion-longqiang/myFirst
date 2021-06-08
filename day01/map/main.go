package main

import (
	"fmt"
	"sort"
)

func main()  {

	m := make(map[int]string)
	m[0] = "ABC"
	fmt.Println(m)
	var m1 map[int]string

	a , ok := m1[1]
	fmt.Println(a ,ok)


	map_1 := map[int]string{4 :"a",5:"B",3:"C",2:"D",1:"R"}
	s := make([]int ,len(map_1))
	i := 0
	for k,_ := range map_1{
		s[i]  = k
		i++
	}
	sort.Ints(s)
	fmt.Println(map_1)


	map_2 := make(map[string]int,len(map_1))
	for k,v := range map_1{
		map_2[v] = k
	}
	fmt.Println(map_1)
	fmt.Println(map_2)


}
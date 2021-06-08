package main

import "fmt"

type person struct{
	name string
	age int
}

type student struct{
	person
	ID string
}
func main()  {
	a := &person{
		"llq",
		22,
	}
	fmt.Println(a)
	a.age = 99
	a.name = "lllll"
	fmt.Println(a)
	A(a)
	fmt.Println(a)
	B(a)
	fmt.Println(a)

	s1 := &student{person:person{"llq",18},ID:"123456"}
	fmt.Println(s1)

	fmt.Println("------------")
	a.print()
	s1.print()
}

func A(a *person) {
	a.name = a.name +"A"
}
func B(b *person) {
	b.name = b.name + "B"
}

func (a *person)print(){
	fmt.Println(a.name , a.age)
}


func (b *student)print(){
	fmt.Println(b.name , b.age,b.ID)
}
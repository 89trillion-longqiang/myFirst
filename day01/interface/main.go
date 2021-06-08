package main

import "fmt"

type Phone interface {
	call()
}

type MyPhone struct{

}
func (p MyPhone)call(){
	fmt.Println("i am MyPhone")
}

func main()  {
	var p Phone
	p = Phone(MyPhone{})
	p  = new(MyPhone)
	p.call()
	disCon(MyPhone{})
	disCon("a")
}

func disCon(inter interface{}) {
	switch inter.(type) {
	case MyPhone: fmt.Println("MyPhone")
	default:
		fmt.Println("Unknow ")
	}
}
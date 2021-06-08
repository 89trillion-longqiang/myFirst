package main

import (
	"errors"
	"fmt"
)

func main()  {

	error_func("llq2")
	fmt.Println("======")
}
func readString(name string) (err error) {
	if name == "llq" {
		return nil
	}else {
		return errors.New("error name")
	}
}
func error_func(name string)  {
	error := readString(name)
	defer func() {
		err := recover()
		if err != nil{
			fmt.Println("error ",25)
		}
	}()
	if error != nil{
		panic("name error")
	}else {
		return
	}
}
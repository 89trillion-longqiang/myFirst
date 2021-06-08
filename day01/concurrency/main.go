package main

import (
	"fmt"
)

func main()  {

	c := make(chan bool,2)
	go Go(c)
	b := <-c
	fmt.Println(b)

	//close(c)


	close(c)
	fmt.Println(cap(c))
}
func Go(c chan bool ){
	fmt.Println("go !")
	c <- true
}
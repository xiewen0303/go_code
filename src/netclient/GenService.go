package netclient


import (
	"fmt"
)

type Test1 struct {
	UserName string
	Age int
	test string
}




func (self *Test1) test2() {
	fmt.Println("world")
}

func (self *Test1) Test1() {
	fmt.Println("world")
}
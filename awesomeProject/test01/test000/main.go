package main

import (
	"fmt"
)
func main() {
	time := "hello goland"
	// str := fmt.Sprintf("\%%s\%",time)
	str := "%"+time+"%"
	fmt.Println(str)

}

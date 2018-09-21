package main

import (
	"bytes"
	"fmt"
)

func main() {

	// Interpret Compare's result by comparing it to zero.
	var a, b []byte
	a=[]byte("B")
	b=[]byte("B")
	if bytes.Compare(a, b) < 0 {
		//a less b
		fmt.Println("a < b")
	}else {
		fmt.Println("a > b")
	}

	fmt.Println(bytes.Equal(a,b))
}

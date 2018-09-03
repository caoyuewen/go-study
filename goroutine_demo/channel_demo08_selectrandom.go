/*
*Desc:4.3 channel --page:187  code:4-14  select
*CreateBy:Cooyw
*Time:2018/9/3
*/
package main

import "fmt"

func main() {
	chanCap := 5
	intChan := make(chan int, chanCap)

	for i := 0; i < chanCap; i++ {
		select {
		case intChan <- 1:
		case intChan <- 2:
		case intChan <- 3:
		}
	}

	for i := 0; i < chanCap; i++ {
		fmt.Printf("%d\n", <-intChan)
	}

}

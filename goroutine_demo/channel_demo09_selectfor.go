/*
*Desc:4.3 channel --page:189  code:4-15  selectfor
*CreateBy:Cooyw
*Time:2018/9/3
*/
package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	close(intChan)

	syncChan := make(chan struct{})

	go func() {

	loop:
		for {
			select {
			case elem, ok := <-intChan:
				if !ok {
					fmt.Println("End.")
					break loop
				}
				fmt.Println("Receive ", elem)
			}
		}
		syncChan <- struct{}{}
	}()

	<-syncChan
}

/*
*Desc:4.3 channel --page:178  code:4-9  channel 的关闭
*CreateBy:Cooyw
*Time:2018/9/2
*/
package main

import "fmt"

func main() {
	dataChan := make(chan int, 5)
	syncC1 := make(chan struct{}, 1)
	syncC2 := make(chan struct{}, 2)

	go func() {
		<-syncC1
		for {
			if elem, ok := <-dataChan; ok {
				fmt.Printf("Received:%d [receiver]\n", elem)
			} else {
				break
			}
		}
		fmt.Println("Done. [receiver]")
		syncC2 <- struct{}{}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			dataChan <- i
			fmt.Printf("Sent: %d [sender]\n", i)
		}

		close(dataChan)
		syncC1 <- struct{}{}
		fmt.Println("Done. [sender]")
		syncC2 <- struct{}{}
	}()

	<-syncC2
	<-syncC2

}

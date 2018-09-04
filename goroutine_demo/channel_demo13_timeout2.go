/*
*Desc:4. 3 channel 195 复用的timer code 4-19
*CreateBy:Cooyw
*Time:2018/9/4
*/
package main

import (
	"time"
	"fmt"
)

func main() {

	intChan := make(chan int, 1)

	go func() {
		for i := 0; i < 5; i++ {
			intChan <- i
			time.Sleep(time.Millisecond * 1000)
		}
		fmt.Println("Sent Done!")
		close(intChan)
	}()

	timeOut := time.Millisecond * 500
	var timer *time.Timer

//loop:
	for {
		if timer == nil {
			timer = time.NewTimer(timeOut)
		} else {
			timer.Reset(timeOut)
		}
		select {
		case elem, ok := <-intChan:
			{
				if !ok {
					fmt.Println("Receive Over!")
					//break loop
					return
				}
				fmt.Println("Receive:", elem)
			}
		case <-timer.C:
			{
				fmt.Println("TimeOut!")
			}
		}
		time.Sleep(timeOut)
	}

}

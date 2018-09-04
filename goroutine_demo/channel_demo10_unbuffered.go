/*
*Desc:4.3 channel --page:189  code:4-16  unbuffered(无缓冲channel)
*CreateBy:Cooyw
*Time:2018/9/3
*/
package main

import (
	"time"
	"fmt"
)

func main() {
	sendInterval := time.Second

	receiveInterval := time.Second * 2
	//receiveInterval := time.Second * 4

	intChan := make(chan int)
	//intChan := make(chan int,5)

	//send
	go func() {
		var ts0, ts1 int64
		for i := 0; i < 5; i++ {
			intChan <- i
			ts1=time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("Sent:", i)
			} else {
				fmt.Printf("Sent:%d [interval:%d s]\n", i, ts1-ts0)
			}
			ts0 = time.Now().Unix()
			time.Sleep(sendInterval)
		}
		close(intChan)
	}()

	var ts0, ts1 int64
loop:
	for {
		select {
		case elem, ok := <-intChan:
			{
				if !ok {
					break loop
				}
				ts1=time.Now().Unix()
				if ts0 == 0 {
					fmt.Println("Received:", elem)
				} else {
					fmt.Printf("Received:%d [interval:%d s]\n", elem, ts1-ts0)
				}
			}
		}
		ts0 = time.Now().Unix()
		time.Sleep(receiveInterval)
	}
	fmt.Println("End.")
}

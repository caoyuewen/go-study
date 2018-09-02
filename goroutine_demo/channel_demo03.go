/*
*Desc: 4.3 channel --page:173  code:4-8
*CreateBy:Cooyw
*Time:2018/9/2
*/
package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

//var MapChan = make(chan map[string]Counter, 1)
var MapChan = make(chan map[string]*Counter, 1)

func main() {
	syncChan := make(chan struct{}, 2)

	go func() {
		for {
			if elem, ok := <-MapChan; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stoped. [receiver]")
		syncChan <- struct{}{}
	}()

	go func() {
		//countMap := map[string]Counter{
		//	"count":Counter{},
		//}
		countMap := map[string]*Counter{
			"count":&Counter{},
		}

		for i := 0; i < 5; i++ {
			MapChan <- countMap
			time.Sleep(time.Microsecond)
			//fmt.Printf("The count map:%v [sender]\n", countMap["count"].count)
			fmt.Printf("The count map:%v [sender]\n", countMap["count"].count)
		}
		close(MapChan)
		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan
}

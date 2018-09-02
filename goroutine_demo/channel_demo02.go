/*
*Desc: 4.3 channel --page:173  code:4-7
*CreateBy:Cooyw
*Time:2018/9/2
*/
package main

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int, 1)

func main() {
	syncChan := make(chan struct{}, 2)

	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				elem["count"]++
			} else {
				break
			}
		}
		syncChan <- struct{}{}
	}()

	go func() {
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {

			mapChan <- countMap
			time.Sleep(time.Microsecond)
			fmt.Printf("The count map:%v.[sender]\n", countMap)
		}
		fmt.Println("Stoped. [receiver]")
		close(mapChan)
		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan

}

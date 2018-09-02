/*
*Desc:4.3 channel --page:182  code:4-10  单向channel
*CreateBy:Cooyw
*Time:2018/9/2
*/
package main

import (
	"fmt"
	"time"
)

var strC = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go receive(strC, syncChan1, syncChan2)
	go send(strC, syncChan1, syncChan2)

	<-syncChan2
	<-syncChan2

}

func receive(strChan <-chan string, sc1 <-chan struct{}, sc2 chan<- struct{}) {

	<-sc1
	for {
		if elem, ok := <-strChan; ok {
			fmt.Println("Reveiced", elem, "[sender]")
		} else {
			break
		}
	}

	fmt.Println("Stopped. [receiver]")
	sc2 <- struct{}{}
}

func send(strChan chan<- string, sc1, sc2 chan<- struct{}) {
	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		if elem == "c" {
			sc1 <- struct{}{}
			fmt.Println("Sent a sync signal. [sender]")
		}
	}

	fmt.Println("Wait 2 seconds... [sender]")
	time.Sleep(time.Second * 2)
	close(strChan)
	sc2 <- struct{}{}
}

/*
*Desc: 4.3 channel --page:173  code:4-6
*CreateBy:Cooyw
*Time:2018/9/2
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	var strChan = make(chan string, 3)
	var syncChan1 = make(chan struct{}, 1)
	var syncChan2 = make(chan struct{}, 2)

	go receiver(strChan,syncChan1,syncChan2)
	go sender(strChan,syncChan1,syncChan2)

	<-syncChan2
	<-syncChan2
}

//接收方
func receiver(strChan chan string, syncChan1, syncChan2 chan struct{}) {
	<-syncChan1
	fmt.Println("Received a sync single and wait a second...[recieve] ")
	time.Sleep(time.Second)

	//for {
	//	if elem, ok := <-strChan; ok {
	//		fmt.Println("Received:", elem, "[receiver]")
	//	} else {
	//		break
	//	}
	//}

	for elem := range strChan {
		fmt.Println("Received:", elem, "[reciever]")
	}

	fmt.Println("Stopd receiver")
	syncChan2 <- struct{}{}
}

//发送方
func sender(strChan chan string, syncChan1, syncChan2 chan struct{}) {
	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		fmt.Println("Sent:", elem, "[sender]")

		if elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("Sent a sync single. [serder]")
		}
	}

	fmt.Println("wait 2 second... [sender]")
	time.Sleep(time.Second * 2)
	close(strChan)
	syncChan2 <- struct{}{}
}

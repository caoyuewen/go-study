/*
*Desc:
*CreateBy:Cooyw
*Time:2018/9/6
*/
package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {

	var rwm sync.RWMutex

	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("Try to lock the reading...[%d]\n", i)
			rwm.RLock()
			fmt.Printf("Locked the reading.[%d]\n", i)
			time.Sleep(time.Second * 2)
			fmt.Printf("Try to Unlock the reading.[%d]\n", i)
			rwm.RUnlock()
			fmt.Printf("Unlock the reading [%d]\n", i)
		}(i)
	}

	time.Sleep(time.Millisecond * 100)
	fmt.Println("Try to lock for writing...")

	//rwm.RLock()
	rwm.Lock()

	fmt.Println("Locked for the writing...")

}

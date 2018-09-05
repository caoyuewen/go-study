/*
*Desc:4. 3 lock 195  page:235 5-2
*CreateBy:Cooyw
*Time:2018/9/5
*/
package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {

	var mutex sync.Mutex

	fmt.Println("Lock the lock (main)")
	mutex.Lock()

	for i:=0;i<=3 ;i++  {
		go func(i int) {
			fmt.Printf("Lock the lock (g%d)\n",i)
			mutex.Lock()
			fmt.Printf("The lock is locked.(g%d)\n",i)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Unlock the lock (main)")
	mutex.Unlock()
	fmt.Println("The lock is unlocked (main)")
	time.Sleep(time.Second)
}






/*
*Desc:4. 3 lock 195  page-235 5-2
*CreateBy:Cooyw
*Time:2018/9/5
*/
package main

import (
	"sync"
	"fmt"
)

func main() {

	/*
		回复运行时候恐慌 在这里是徒劳  go1.8之前可以恢复，但是会导致很严重的问题(比如这个重复解锁的goroutine会永久阻塞)
	*/
	defer func() {
		fmt.Println("Try to recover the painc")
		if p := recover(); p != nil {
			fmt.Printf("Recovered the painc(%#v).\n", p)
		}
	}()

	var mutex sync.Mutex

	fmt.Println("Lock the lock")
	mutex.Lock()
	fmt.Println("The lock is locked")
	fmt.Println("Unlock the lock")
	mutex.Unlock()
	fmt.Println("The lock is unlocked")
	fmt.Println("Unlock the lock again")
	mutex.Unlock()

}

/*
*Desc:
*CreateBy:Cooyw
*Time:2018/9/9
*/
package main

import (
	"math/rand"
	"fmt"
	"sync"
)

func main() {

	var count int

	//只会执行一次
	var once sync.Once

	max := rand.Intn(100)
	fmt.Println(max)
	for i := 0; i < max; i++ {
		once.Do(func() {
			count++
		})
	}

	fmt.Println(count)

}

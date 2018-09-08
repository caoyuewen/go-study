/*
*Desc:
*CreateBy:Cooyw
*Time:2018/9/8
*/
package main

import (
	"sync/atomic"
	"fmt"
)

func main() {

	var NN int32 = -2

	var i32 int32
	i32 = 20
	//原子操作 增或减 将int32的值增加3
	newi32 := atomic.AddInt32(&i32, 3)
	fmt.Println(newi32)

	//原子操作 增或减 将int32的值减少3
	newi32 = atomic.AddInt32(&i32, -6)
	fmt.Println(newi32)

	//原子操作 减少一个变量的值 不能这样写  newi32 = atomic.AddInt32(&i32, -NN))
	newi32 = atomic.AddInt32(&i32, ^int32(-NN-1))
	fmt.Println(newi32)

	fmt.Println(^int32(newi32))






}

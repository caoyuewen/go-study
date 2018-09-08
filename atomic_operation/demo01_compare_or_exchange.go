/*
*Desc:
*CreateBy:Cooyw
*Time:2018/9/8
*/
package main

import (
	"sync/atomic"
)

func main() {

	//CAS:比较并交换 优势：在不创建互斥量和不形成临界区资源的情况下，完成并发安全的值替换操作

	var delta int32

	addValue(delta)
}

var value int32

func addValue(delta int32) {
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, v+delta) {
			break
		}
	}
}

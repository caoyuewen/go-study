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

	var countVal atomic.Value

	countVal.Store([]int{1, 3, 5, 7})
	//anotherStore(countVal)
	anotherStore(&countVal)

	fmt.Printf("The count value:%v \n", countVal)

}

/*func anotherStore(countVal atomic.Value) {*/
func anotherStore(countVal *atomic.Value) {

	countVal.Store([]int{2, 4, 6, 8})
}

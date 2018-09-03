/*
*Desc:4.3 channel --page:183  code:4-11  chanconv
*CreateBy:Cooyw
*Time:2018/9/3
*/
package main

import "fmt"

func main() {


	//chan interface{} 、chan<- interface{} 、<-chan interface{} 不是同一类型


	var ok bool
	ch := make(chan int, 1)

	//判断ch 是不是单向chan(只能接收)
	_, ok = interface{}(ch).(<-chan int)
	fmt.Println("chan int => <-chan int:", ok)

	_, ok = interface{}(ch).(chan<- int)
	fmt.Println("chan int => <-chan int:", ok)

	sch := make(chan<- int)
	_, ok = interface{}(sch).(chan int)
	fmt.Println("chan<- int => chan int:", ok)

	rch := make(<-chan int)
	_, ok = interface{}(rch).(chan int)
	fmt.Println("chan<- int => chan int:", ok)

}

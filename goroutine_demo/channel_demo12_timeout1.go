/*
*Desc:
*4. 3 channel 195 code 4-18
我并发地在in tchan通道进行发送和接收操作。发送操作的延时是1s。接收操作没
有延时，但是有对操作超时的设定。关键在于select语句中的第二个case表达式，这里
初始化了一个相对到期时间为500ms的定时器，并试图立即从它的字段C中接收元素值。
一旦定时器到期，该接收操作就会完成，select语句的执行也就结束了。此时发送操作
还未进行，因此第一个case失去了被选中的机会。如此就实现了操作超时。
你可能会觉得time. Newfimer(time. Millisecond s00). C太烦琐了。这样的话，你可以
用time. After(time. Millisecond 50)春换之。它与前者是等价的，都可以表示经转换的
通知通道。time. Fter函数会新建一个定时器，并把它的字段C作为结果返回。此函数的
作用相当简单，即:对超时的设定提供了一种快捷方式。
如前文所述，select语句与for语句连用可以持续地从一个通道接收元素值。但是，
若每次接收时都初始化一个定时器显然有些浪费，好在定时器是可以复用的。
*CreateBy:Cooyw
*Time:2018/9/3
*/
package main

import (
	"time"
	"fmt"
)

func main() {

	intChan := make(chan int)
	
	go func() {
		time.Sleep(time.Second)
		intChan <- 1
	}()

	select {
	case elem := <-intChan:
		fmt.Printf("Received:%d\n", elem)

		//case <-time.NewTimer(time.Millisecond * 500).C:
	case <-time.After(time.Millisecond * 500):

		fmt.Println("Timeout!")
	}

}

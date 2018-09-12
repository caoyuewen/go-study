/*
*Desc:list包实现了双向链表
*CreateBy:Cooyw
*Time:2018/9/12
*/
package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()

	e4 := l.PushBack(4)
	e1 := l.PushFront(1)

	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	//遍历链表
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)

	}
}

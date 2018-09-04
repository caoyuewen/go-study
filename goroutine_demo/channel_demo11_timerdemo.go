/*
*Desc: 4.3 channel --page:194  code:4-17  time包的使用
*CreateBy:Cooyw
*Time:2018/9/3
*/
package main

import (
	"time"
	"fmt"
)

func main() {

	timer := time.NewTimer(2 * time.Second)
	fmt.Printf("Persent time:%v\n", time.Now())
	//定时器阻塞2秒后执行
	expirationTime := <-timer.C
	fmt.Printf("Expiration time:%v\n",expirationTime)
	fmt.Printf("Stop timer :%v.\n",timer.Stop())

}


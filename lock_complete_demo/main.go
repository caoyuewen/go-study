/*
*Desc:
*CreateBy:Cooyw
*Time:2018/9/6
*/
package main

import (
	mdf "github.com/caoyuewen/go-study/lock_complete_demo/mydatafile"
	"fmt"
	"os"
)

func main() {
	var path = "F:/gocode/src/github.com/caoyuewen/go-study/lock_complete_demo/mydatafile/txt.txt"
	file, err := os.Open(path)
	n, err := file.WriteString("你好")
	if err!=nil {
		fmt.Println("main-err:",err.Error())
		return
	}
	if err != nil {
		fmt.Println("main-err:",err.Error())
		return
	}

	bytes:=make([]byte,1024)
	n, err = file.Read(bytes)
	if err!=nil {
		fmt.Println("main-err:",err.Error())
		return
	}

	fmt.Println(string(bytes))
	fmt.Println(n)

}

func myDataFileTest() {
	var path = "F:/gocode/src/github.com/caoyuewen/go-study/lock_complete_demo/mydatafile/txt.txt"

	file, err := mdf.NewDataFile(path, 512)
	if err != nil {
		fmt.Println("main->err:", err.Error())
		return
	}

	wsn, err := file.Write([]byte("hello!"))
	if err != nil {
		fmt.Println("main->err:", err.Error())
		return
	}
	fmt.Println(wsn)

	wsn, err = file.Write([]byte("hello2!"))
	if err != nil {
		fmt.Println("main->err:", err.Error())
		return
	}

	fmt.Println(wsn)

	wsn, err = file.Write([]byte("这是第三次写入!"))
	if err != nil {
		fmt.Println("main->err:", err.Error())
		return
	}

	fmt.Println(wsn)

	fmt.Println(file.DataLen())

	rsn, d, err := file.Read()
	if err != nil {
		fmt.Println("main->err:", err.Error())
	}
	fmt.Println(string(d))

	fmt.Println(rsn)
}

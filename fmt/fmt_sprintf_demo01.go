/**
 * Description: 
 * User: 1067
 * Date: 2018-09-14
 * Time: 17:56
 */

package main

import "fmt"

func main() {
	//fmt.Sprintf可以做进制转换
	/*
	%d	表示为十进制
	%o	表示为八进制
	%x	表示为十六进制，使用a-f
	%X	表示为十六进制，使用A-F
	*/
	i10 := 32                     //10进制的32
	i2 := fmt.Sprintf("%b", i10)  //二进制
	i8 := fmt.Sprintf("%o", i10)  //八进制
	i16 := fmt.Sprintf("%x", i10) //16进制 小写

	fmt.Println(i2)
	fmt.Println(i8)
	fmt.Println(i10)
	fmt.Println(i16)

}

/*
*Desc:4.3 channel --page:187  code:4-13  select
*CreateBy:Cooyw
*Time:2018/9/3
*/
package main

import "fmt"

var intChan1 chan int
var intChan2 chan int
var channels = []chan int{intChan1, intChan2}

var numbers = []int{1, 2, 3, 4, 5}

func main() {

	select {
	case getChan(0) <- getNumber(0):
		fmt.Println("1th case is selected")
	case getChan(1) <- getNumber(1):
		fmt.Println("2th case is selected")
	default:
		fmt.Println("Defautl!")
	}
	//fmt.Println(<-channels[0])
	//fmt.Println(<-channels[1])
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}

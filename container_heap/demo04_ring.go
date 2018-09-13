/**
 * Description:
	ring实现了环形链表的操作。
 * User: 1067
 * Date: 2018-09-13
 * Time: 9:34
 */

package main

import (
	"container/ring"
	"fmt"
)

func main() {

	const rLen = 3

	//创建新的ring
	r := ring.New(rLen)

	//为ring赋值
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	printRing := func(v interface{}) {

		fmt.Print(v, " ")
	}

	r.Do(printRing) // 0 1 2
	fmt.Println()

	//将r之后的第二个元素的值*2
	r.Move(2).Value = r.Move(2).Value.(int) * 2

	r.Do(printRing) // 0 1 4
	fmt.Println()

	//删除r 到 r+2之间的元素 即删除 r2 即r+2
	//返回的是被删除掉的元素组成的ring指针
	result := r.Link(r.Move(2))

	r.Do(printRing) //0 4
	fmt.Println()

	fmt.Println(result.Value) //1

	//创建另一个ring
	another := ring.New(rLen)
	another.Value = 7
	another.Next().Value = 8 // 给 another + 1 表示的元素赋值，即第二个元素
	another.Prev().Value = 9 // 给 another - 1 表示的元素赋值，即第三个元素

	another.Do(printRing) //7 8 9
	fmt.Println()

	// 插入another到r后面，返回插入前r的下一个元素
	result = r.Link(another)

	r.Do(printRing) //0 7 8 9 4
	fmt.Println()

	result.Do(printRing) //4 0 7 8 9
	fmt.Println()

	// 删除r之后的三个元素，返回被删除元素组成的Ring的指针
	result = r.Unlink(3)

	r.Do(printRing)//0 4
	fmt.Println()

	result.Do(printRing)//7 8 9
	fmt.Println()


}

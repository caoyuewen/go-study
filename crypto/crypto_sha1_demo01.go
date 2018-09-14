/**
 * Description: 
 * User: 1067
 * Date: 2018-09-13
 * Time: 15:21
 */

package main

import (
	"crypto/sha1"
	"fmt"
)
//12=32356a192b7913b04c54574d18c28d46e6395428ab
//1 =
func main() {
	s := "1654654"
	h := sha1.New()

	h.Write([]byte(s))

	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要
	bs := h.Sum(nil)

	//SHA1 值经常以 16 进制输出，例如在 git commit 中。使用%x 来将散列结果格式化为 16 进制字符串。
	fmt.Println(s)
	fmt.Printf("%x\n", bs)

}

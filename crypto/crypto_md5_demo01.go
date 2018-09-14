/**
 * Description: 
 * User: 1067
 * Date: 2018-09-14
 * Time: 15:40
 */

package main

import (
	"crypto/md5"
	"fmt"
	"crypto/sha1"
)

func main() {

	/*
	MD5不可逆（无法从MD5还原为原来 的信息）
	但是MD5可以被伪造A->MD5A,可以伪造 信息B->MD5A
	A,B的MD5一样

	始终返回固定长度的字符
	*/

	data := []byte("this is something important code should encrypt")

	sum := md5.Sum(data)
	md5Code := fmt.Sprintf("%x", sum)
	fmt.Println(md5Code)

	//###########################################

	hash := md5.New()
	hashBytes := hash.Sum(data)
	hashCode := fmt.Sprintf("%x", hashBytes)
	fmt.Println(hashCode)

	//#################################

	shash := sha1.New()
	shashBytes := shash.Sum(data)
	shashCode := fmt.Sprintf("%x", shashBytes)
	fmt.Println(shashCode)

}

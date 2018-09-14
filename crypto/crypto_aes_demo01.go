/**
 * Description: 
 * User: 1067
 * Date: 2018-09-14
 * Time: 10:31
 */

package main

import (
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)
//93e95aed25bf8f997378d5f61546458af81016bf06ba5b7e639a92dc7c77aa3e5a1f0190d6f359c6
//155ac4d88bc5de78a0631842ad7a6dcc33a9b1c5bbe37e7e08461a524ab28d0265038c9e0610814e
func main() {

	nonce := "37b8e8a308c354048d245f6d"
	//key := "AES256Key-32Characters1234567890"
	key := "4f5f8d7a9x5s6g9a"	//秘钥长度需要时AES-128(16bytes)或者AES-256(32bytes)
	plainText := "我是要传输的内容"

	cipherText := ExampleNewGCM_encrypt(plainText, key, nonce)
	newPlain := ExampleNewGCM_decrypt(cipherText, key, nonce)

	fmt.Println("plain:", plainText)
	fmt.Println("cipher:", cipherText)
	fmt.Println("new plain:", newPlain)
}

//解密
func ExampleNewGCM_encrypt(src, k, n string) string {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte(k)
	plainText := []byte(src)
	block, err := aes.NewCipher(key)//生成加密用的block
	if err != nil {
		panic(err)
	}
	nonce, _ := hex.DecodeString(n)

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, plainText, nil)
	return fmt.Sprintf("%x", ciphertext)

}

//加密后生成密文
func ExampleNewGCM_decrypt(src, k, n string) string {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte(k)
	cipherText, _ := hex.DecodeString(src)
	nonce, _ := hex.DecodeString(n)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	asegcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	plaintext, err := asegcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}

	return string(plaintext)

}




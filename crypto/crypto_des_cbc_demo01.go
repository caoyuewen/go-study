/**
 * Description: 
 * User: 1067
 * Date: 2018-09-14
 * Time: 11:31
 */

package main

import (
	"crypto/des"
	"bytes"
	"crypto/cipher"
	"fmt"
	"encoding/hex"
)

func main() {

	data := "我是传输数据"
	//key := "4f5f8d7a9x5s6g9a"
	key := "84848484"
	nonce := "98798797"

	dataDesCode := EncryptDES_CBC(data, key,nonce)
	dataDexEnc := DecryptDES_CBC(dataDesCode, key,nonce)

	fmt.Println(data)
	fmt.Println(dataDesCode)
	fmt.Println(dataDexEnc)

}

//CBC加密
func EncryptDES_CBC(src, key, nonce string) string {
	data := []byte(src)
	keyByte := []byte(key)
	nonceByte := []byte(nonce)

	block, err := des.NewCipher(keyByte) //key的长度一定要是8个字节
	if err != nil {
		panic(err)
	}

	data = PKCS5Padding(data, block.BlockSize())
	//获取CBC加密模式
	//iv := keyByte //用密钥作为向量(不建议这样使用)
	iv := nonceByte
	mode := cipher.NewCBCEncrypter(block, iv)
	out := make([]byte, len(data))
	mode.CryptBlocks(out, data)
	return fmt.Sprintf("%X", out)
}

//CBC解密
func DecryptDES_CBC(src, key, nonce string) string {
	data, _ := hex.DecodeString(src)
	keyByte := []byte(key)
	nonceByte := []byte(nonce)

	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}

	//iv := keyByte //用密钥作为向量(不建议这样使用)
	iv := nonceByte
	mode := cipher.NewCBCDecrypter(block, iv)

	plaintext := make([]byte, len(data))
	mode.CryptBlocks(plaintext, data)
	plaintext = PKCS5UnPadding(plaintext)
	//return fmt.Sprintf("%X", plainText)
	return string(plaintext)
}



//明文补码算法
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//明文减码算法
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

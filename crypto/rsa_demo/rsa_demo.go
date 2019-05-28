package main

import (
	"encoding/pem"
	"io/ioutil"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"github.com/micro/go-micro/errors"
	"fmt"
)

const (
	PriKey = "private.pem"
	PubKey = "public.pem"
)

//加密
func RsaEncrypt(origData []byte) (data []byte, err error) {

	pubKey, err := ioutil.ReadFile(PubKey)
	if err != nil {
		return
	}

	//解密pem格式的公钥
	block, _ := pem.Decode(pubKey)

	//解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	//pub, err := x509.ParsePKCS1PublicKey(block.Bytes)

	//类型断言
	pub := pubInterface.(*rsa.PublicKey)

	data, err = rsa.EncryptPKCS1v15(rand.Reader, pub, origData)

	return
}

//解密
func RsaDecrypt(cipherText []byte) (data []byte, err error) {
	priKey, err := ioutil.ReadFile(PriKey)
	if err != nil {
		return
	}

	//解密
	block, _ := pem.Decode(priKey)
	if block == nil {
		err = errors.New("", "private key error!", -1)
		return
	}

	//解析PKCS1格式的私钥
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	data, err = rsa.DecryptPKCS1v15(rand.Reader, key, cipherText)
	return
}

func main() {

	origData := "你好我是"

	fmt.Printf("原始数据:\n%s\n",origData)

	//加密
	cipherText, err := RsaEncrypt([]byte(origData))
	if err != nil {
		panic(err)
	}

	fmt.Printf("加密后的数据:\n%s\n",string(cipherText))

	//解密
	data, err := RsaDecrypt(cipherText)
	if err != nil {
		panic(err)
	}

	fmt.Printf("解密后的数据:\n%s\n",string(data))

}

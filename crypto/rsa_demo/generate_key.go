//生成rsa密钥对

package main

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
)

const (
	rsaKeyBits        = 2048          //模值长度
	rsaPriKeyFileName = "private.pem" //生成的私钥文件名
	rsaPubKeyFileName = "public.pem"  //生成的公钥文件名
)

func main() {
	err := GenRsaKey(rsaKeyBits)
	if err != nil {
		panic(err)
	}
}

//RSA公钥私钥产生  PKCS1
func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derPriKeyStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derPriKeyStream,
	}
	file, err := os.Create(rsaPriKeyFileName)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPubKeyStream, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derPubKeyStream,
	}
	file, err = os.Create(rsaPubKeyFileName)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

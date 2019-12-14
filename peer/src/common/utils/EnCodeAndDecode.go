package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	_ "encoding/base64"
	"encoding/pem"
	"errors"
	_ "flag"
	"fmt"
	"io/ioutil"
	"os"
)

var decrypted string
var PiKey, PuKey []byte

func InitLoadKeyAndValue() {

	var err error
	//读公钥
	PuKey, err = ioutil.ReadFile(PUBLIC_KEY_PATH)
	if err != nil {
		os.Exit(-1)
	}
	//读私钥
	PiKey, err = ioutil.ReadFile(PRIVATE_KEY_PATH)
	if err != nil {
		os.Exit(-1)
	}
}

func main() {
	var data []byte
	var err error

	//数据加密
	data, err = RsaEncrypt([]byte("fyxichen"))
	fmt.Println(data)
	if err != nil {
		panic(err)
	}
	//数据解密
	origData, err := RsaDecrypt(data)
	fmt.Println(origData)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}

// 加密
/*
	block.Bytes:  公钥
	block.Headers: 空
	block.Type:  "公钥"
*/
func RsaEncrypt(origData []byte) ([]byte, error) {

	block, _ := pem.Decode(PuKey)

	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(PiKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

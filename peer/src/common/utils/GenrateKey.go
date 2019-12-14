package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

//D:/GoProjects/DairyOriginateSystem/peer/src/common/keyPair/private.pem//
const bits int = 1024
const PRIVATE_KEY_PATH = "./common/utils/keyPair/private.pem"
const PUBLIC_KEY_PATH = "./common/utils/keyPair/public.pem"

func init() {

	flag1, _ := PathExists(PRIVATE_KEY_PATH)
	flag2, _ := PathExists(PUBLIC_KEY_PATH)
	if flag1 || flag2 {
		fmt.Println("密钥文件存在.无需生成")
	} else {
		if err := GenRsaKey(); err != nil {
			fmt.Println(err.Error())
			log.Fatal("密钥文件生成失败！")

		}
		log.Println("密钥文件生成成功！")
	}
}

func GenRsaKey() error {
	//生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "私钥",
		Bytes: derStream,
	}
	fmt.Println("1")
	file, err := os.Create(PRIVATE_KEY_PATH)
	fmt.Println("2")
	if err != nil {

		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	//生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "公钥",
		Bytes: derPkix,
	}

	file, err = os.Create(PUBLIC_KEY_PATH)

	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

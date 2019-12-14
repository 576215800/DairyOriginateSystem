package main

import (
	"common/utils"
	"config"
	"core"
	"fmt"
	"logging"
	_ "testFile"
	"web"
)

func main() {
	utils.InitLoadKeyAndValue()
	test := "hello world"
	data, _ := utils.RsaEncrypt([]byte(test))
	fmt.Println(data)
	ori_data, _ := utils.RsaDecrypt(data)
	fmt.Println(ori_data)
	fmt.Println(string(ori_data))
	fmt.Println()
	core.NewBlockchain()
	config.Initialize()
	logging.Initialize()
	web.Initialize()

}

package utils

import (
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path) //判断文件是否存在
	if err == nil {
		return true, nil
	} else {
		return false, err
	}
}

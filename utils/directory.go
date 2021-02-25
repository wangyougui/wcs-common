package utils

import "os"

// @description   文件目录是否存在
// @auth                     （2021/02/25  20:22）
// @param     path            string
// @return    err             error
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

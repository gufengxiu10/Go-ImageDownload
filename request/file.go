package request

import "os"

//判断文件是否存在
func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return os.IsExist(err)
}

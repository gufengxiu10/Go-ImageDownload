package request

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func download(url string, path string) {
	fileNameSplit := strings.Split(url, "/")
	fileName := fileNameSplit[len(fileNameSplit)-1]
	fileName = path + "/" + fileName

	if FileExists(fileName) {
		return
	}

	response, err := http.Get(url)
	if err != nil {
		println(err)
		return
	}
	file, _ := os.Create(fileName)
	defer file.Close()
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	file.Write(body)
}

func downLoadAll(url []string, path string) {
	for _, v := range url {
		download(v, path)
	}
}

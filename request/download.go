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
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	//增加header选项
	reqest.Header.Add("referer", "https://pixiviz.pwp.app/")
	reqest.Header.Add("origin", "https://pixiviz.pwp.app")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 Edg/95.0.1020.44")
	if err != nil {
		println(err)
		return
	}
	response, _ := client.Do(reqest)
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

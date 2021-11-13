package request

import (
	"io"
	"net/http"
	"os"
	"strings"
)

type option struct {
	r http.RoundTripper
}

func (b option) RoundTrip(reqest *http.Request) (*http.Response, error) {
	reqest.Header.Add("referer", "https://pixiviz.pwp.app/")
	reqest.Header.Add("origin", "https://pixiviz.pwp.app")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 Edg/95.0.1020.44")
	return b.r.RoundTrip(reqest)
}

func download(url string, path string) {
	fileNameSplit := strings.Split(url, "/")
	fileName := fileNameSplit[len(fileNameSplit)-1]
	fileName = path + "/" + fileName

	if FileExists(fileName) {
		return
	}
	client := &http.Client{
		Transport: option{r: http.DefaultTransport},
	}
	response, err := client.Get(url)
	//增加header选项
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

package img

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type pixiviz struct {
	roundTripper http.RoundTripper
	header       map[string]string
}

//判断文件是否存在
func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return os.IsExist(err)
}

func CreatePix() *pixiviz {
	var header map[string]string = make(map[string]string)
	header["Referer"] = "https://pixiviz.pwp.app/"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 Edg/95.0.1020.44"

	return &pixiviz{
		header:       header,
		roundTripper: http.DefaultTransport,
	}
}

func (img pixiviz) RoundTrip(request *http.Request) (*http.Response, error) {
	for index, value := range img.header {
		request.Header.Add(index, value)
	}

	return img.roundTripper.RoundTrip(request)
}

func (img *pixiviz) Rand() {
	response, _ := img.send("https://pixiviz.pwp.app/api/v1/illust/rank?mode=day&date=2021-11-14&page=1", "get")

	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	img.getUrl(body)
}

func (img *pixiviz) send(url string, method string) (*http.Response, error) {

	client := &http.Client{
		Transport: img,
	}

	request, err := http.NewRequest(strings.ToUpper(method), url, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(request)
}

func (img *pixiviz) getUrl(body []byte) {
	doc := make(map[string][]map[string]interface{})
	json.Unmarshal(body, &doc)
	// urls := make([]string, 0)
	for _, v := range doc["illusts"] {

		if len(v["meta_pages"].([]interface{})) == 0 {
			urlInterface := v["meta_single_page"].(map[string]interface{})
			url := urlInterface["original_image_url"].(string)
			url = strings.Replace(url, "i.pximg.net", "pixiv-image-lv.pwp.link", 1)
			path := "image/pixiviz/rank/"
			os.MkdirAll(path, 0775)
			img.download(url, path)
		} else {
			imageUrls := v["meta_pages"].([]interface{})
			path := "image/pixiviz/rank/" + strconv.Itoa(int(v["id"].(float64)))
			os.MkdirAll(path, 0775)
			for _, imageV := range imageUrls {
				url := imageV.(map[string]interface{})["image_urls"].(map[string]interface{})["original"]
				url = strings.Replace(url.(string), "i.pximg.net", "pixiv-image-lv.pwp.link", 1)
				img.download(url.(string), path)
			}
		}
	}

}

func (img *pixiviz) download(url string, path string) {

	fileNameSplit := strings.Split(url, "/")
	fileName := fileNameSplit[len(fileNameSplit)-1]
	fileName = path + "/" + fileName

	if FileExists(fileName) {
		return
	}

	client := &http.Client{
		Transport: img,
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
	fmt.Println("下载完成")
}

func (img *pixiviz) DownLoadAll(url []string, path string) {
	for _, v := range url {
		img.download(v, path)
	}
}

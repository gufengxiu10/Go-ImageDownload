package img

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type pixiviz struct {
	roundTripper http.RoundTripper
	header       map[string]string
}

func CreatePix() *pixiviz {
	var header map[string]string = make(map[string]string)
	header["Referer"] = "https://pixiviz.pwp.app/"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 Edg/95.0.1020.44"

	return &pixiviz{
		header: header,
	}
}

func (img pixiviz) RoundTrip(request *http.Request) (*http.Response, error) {
	for index, value := range img.header {
		request.Header.Add(index, value)
	}
	return img.roundTripper.RoundTrip(request)
}

func (img *pixiviz) Rand() {
	response, _ := img.send("https://pixiviz-api-us.pwp.link/v1/illust/rank?mode=day&date=2021-11-07&page=1", "Get")

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

func (img *pixiviz) send(url string, method string) (*http.Response, error) {

	img.roundTripper = http.DefaultTransport
	client := &http.Client{
		Transport: img,
	}

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(request)
}

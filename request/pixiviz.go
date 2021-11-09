package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//https://pixiviz-api-us.pwp.link/v1/illust/rank?mode=day&date=2021-11-07&page=18
type PixivizInterface interface {
	Rand()
}

type PixivizStruct struct {
}

type MyRoundTripper struct {
	r http.RoundTripper
}

type illusts struct {
	meta_pages []map[string]map[string]string
}

func (mrt MyRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Referer", "https://pixiviz.pwp.app/")
	r.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 Edg/95.0.1020.44")
	return mrt.r.RoundTrip(r)
}

func (b *PixivizStruct) Rand() {
	url := "https://pixiviz-api-us.pwp.link/v1/illust/rank?mode=day&date=2021-11-07&page=1"
	Client := &http.Client{
		Transport: MyRoundTripper{r: http.DefaultTransport},
	}

	response, err := Client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	var tempMap = make(map[string][]map[string]interface{})

	err = json.Unmarshal([]byte(string(body)), &tempMap)

	if err != nil {
		panic(err)
	}

	for _, v := range tempMap["illusts"] {

		fmt.Printf("%T\n", v)
		// var c []interface{} = v["meta_pages"]
		// fmt.Println(page)
	}
}

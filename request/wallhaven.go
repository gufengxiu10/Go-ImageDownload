package request

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type WallhavenApi interface {
	Random()
	Search(keyWord string)
}

type WallhavenSend struct {
}

func (b *WallhavenSend) Random() {

	path := "image/wallhaven/random/" + time.Now().Format("2006-01-02")
	os.MkdirAll(path, 0775)

	client, err := http.Get("https://wallhaven.cc/random?seed=HY871P")
	if err != nil {
		log.Fatal(err)
	}

	defer client.Body.Close()
	doc, err := goquery.NewDocumentFromReader(client.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".thumb-listing-page figure").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find("img").Attr("data-src")
		urlMap := strings.Split(url, "/")
		currentUrl := urlMap[len(urlMap)-1]
		if s.Find(".png").Text() != "" {
			currentUrl = strings.Replace(currentUrl, "jpg", "png", 1)
		}

		allUrl := b.join(currentUrl)
		download(allUrl, path)
	})
}

func (b *WallhavenSend) Search(key string) {

	path := "image/wallhaven/search/" + key
	os.MkdirAll(path, 0775)

	client, err := http.Get("https://wallhaven.cc/search?q=" + key + "&page=1")
	if err != nil {
		log.Fatal(err)
	}

	defer client.Body.Close()
	doc, err := goquery.NewDocumentFromReader(client.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".thumb-listing-page figure").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find("img").Attr("data-src")
		urlMap := strings.Split(url, "/")
		currentUrl := urlMap[len(urlMap)-1]
		if s.Find(".png").Text() != "" {
			currentUrl = strings.Replace(currentUrl, "jpg", "png", 1)
		}

		allUrl := b.join(currentUrl)
		download(allUrl, path)
	})
}

func (b *WallhavenSend) prefix(url string) string {
	return fmt.Sprintf("%c", url[0]) + fmt.Sprintf("%c", url[1])
}

func (b *WallhavenSend) join(url string) string {
	allUrlSlice := []string{"https://w.wallhaven.cc/full/", b.prefix(url), "/wallhaven-", url}
	allUrl := strings.Join(allUrlSlice, "")
	return allUrl
}

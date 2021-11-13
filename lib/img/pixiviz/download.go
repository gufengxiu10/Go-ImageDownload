package pixiviz

import "net/http"

type option struct {
	t http.RoundTripper
}

type client interface {
	Get()
}

type clienStruct struct {

}

func(t option) RoundTrip(request *http.Request) (*http.Response,error){
	request.Header.Add("Referer", "https://pixiviz.pwp.app/")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 Edg/95.0.1020.44")
	return t.t.RoundTrip(request)
}


func(b clienStruct) rand(){
}

client := &http.Client{
	Transport : option{t:http.DefaultTransport}
}

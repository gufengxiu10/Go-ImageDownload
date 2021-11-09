package main

import "img/request"

func main() {
	var b request.PixivizInterface = new(request.PixivizStruct)
	b.Rand()

	// client := &http.Client{}

	// for _, v := range url {
	// 	fileNameSplit := strings.Split(v, "/")
	// 	fileName := fileNameSplit[len(fileNameSplit)-1]
	// 	_, err := os.Stat(fileName)
	// 	if err == nil {
	// 		continue
	// 	}

	// 	res, err := client.Get(v)
	// 	if err != nil {
	// 		continue
	// 	}

	// 	defer res.Body.Close()

	// 	body, _ := io.ReadAll(res.Body)
	// 	img, _ := os.Create(fileName)
	// 	defer img.Close()
	// 	img.Write(body)
	// }
}

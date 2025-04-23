package http

import(
	"net/http"
	"io/ioutil"
	"bytes"

	"longchain.com/memoriae/orePool/log"
)

func Get(url string) string {
	resp, err := http.Get(url)
	println(url)
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	return string(body)
}
// post application/json
func PostByJson(url,data string) string {
	body_type := "application/json;charset=utf-8"
	req := bytes.NewBuffer([]byte(data))
	resp, err := http.Post(url, body_type, req)
	if err != nil {
		log.Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	return string(body)
}
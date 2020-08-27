package httpwasm

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string) (result string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("js.fetch:mode", "cors")
	if err != nil {
		log.Println(err)
		return
	}
	return process(req)
}

func Post(url string, body []byte) (result string, err error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	req.Header.Add("js.fetch:mode", "cors")
	if err != nil {
		log.Println(err)
		return
	}
	return process(req)
}

func MiscRequest(url, method string, body []byte) (result string, err error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	req.Header.Add("js.fetch:mode", "cors")
	if err != nil {
		log.Println(err)
		return
	}
	return process(req)
}

func process(req *http.Request) (result string, err error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	result = string(b)
	return
}

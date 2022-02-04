package restapi

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var Proxy string = ""

type HttpClient struct {
	Client *http.Client
}

func (r *HttpClient) Get(url string) ([]byte, error) {
	resp, err := r.Client.Get(url)
	if err != nil {
		return []byte{}, err
	}
	dat, _ := ioutil.ReadAll(resp.Body)
	return dat, nil
}

func (r *HttpClient) Post(url string, dat []byte) (bool, error) {
	url = url + "?overwrite=true"
	resp, err := r.Client.Post(url, "application/octet-stream", bytes.NewBuffer(dat))
	if err != nil {
		return false, err
	}
	if resp.StatusCode == 200 {
		return true, nil
	} else {
		dat, _ := ioutil.ReadAll(resp.Body)
		log.Printf("%v\n", string(dat))
	}
	return false, nil
}

func (r *HttpClient) PostAppend(url string, dat []byte) (bool, error) {
	resp, err := r.Client.Post(url, "application/octet-stream", bytes.NewBuffer(dat))
	if err != nil {
		return false, err
	}
	if resp.StatusCode == 200 {
		return true, nil
	} else {
		dat, _ := ioutil.ReadAll(resp.Body)
		log.Printf("%v\n", string(dat))
	}
	return false, nil
}

func (r *HttpClient) StatusCode(url string) (int, error) {
	resp, err := r.Client.Get(url)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}

func NewHttpClient() HttpClient {
	c := http.Client{
		Timeout:       time.Duration(10) * time.Second,
		CheckRedirect: nil,
	}
	if Proxy != "" {
		url, _ := url.Parse(Proxy)
		purl := http.ProxyURL(url)
		c.Transport = &http.Transport{
			Proxy: purl,
		}
	}

	rc := HttpClient{Client: &c}

	return rc

}

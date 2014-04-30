package helper

import (
	//	"fmt"
	"bytes"
	"image"
	"image/jpeg"
	//"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HttpHelper struct {
	RequestURL *url.URL
	PostData   url.Values
}

func (this *HttpHelper) GetImage() (image.Image, error) {
	b, err, _ := this.Get()
	if err != nil {
		return nil, err
	}
	img, err2 := jpeg.Decode(bytes.NewReader(b))
	return img, err2
}

func (this *HttpHelper) Get() (databytes []byte, err error, status int) {
	resp, err := http.Get(this.RequestURL.String())
	defer resp.Body.Close()
	if err != nil {
		//log.Fatal(err)
		if resp == nil {
			return nil, err, 0
		} else {
			return nil, err, resp.StatusCode
		}
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if resp == nil {
			return nil, err, 0
		} else {
			return nil, err, resp.StatusCode
		}
	}
	//fmt.Printf("body length: %d \r\n", len(bytes))
	//ioutil.WriteFile("test.txt", bytes, 777)
	return bytes, nil, resp.StatusCode
}

package utils

import (
	"aos/pkg/setting"
	"time"

	"github.com/imroc/req"
)

type HttpClient struct {
	// HttpClientHandle *req.Req
	Debug bool
}

var HttpHandle = initHttpClient()

func initHttpClient() *HttpClient {
	var httpClient = new(HttpClient)
	// httpClient.HttpClientHandle = req.New()
	httpClient.Debug = false
	req.SetTimeout(5 * time.Second)
	return httpClient
}

func (hc *HttpClient) handle(paramData map[string]interface{}, headerParam map[string]string) (req.Param, req.Header) {
	header := req.Header{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	if headerParam != nil {
		for k, v := range headerParam {
			header[k] = v
		}
	}
	if hc.Debug {
		req.Debug = true
	}
	param := req.Param(paramData)
	return param, headerParam
}

func (hc *HttpClient) Post(url string, paramData map[string]interface{}, headerParam map[string]string) (interface{}, error) {
	param, header := hc.handle(paramData, headerParam)
	r, err := req.Post(url, param, header)
	if err != nil {
		setting.Logger.Infof("http:request:"+url+"===error", err)
	}
	var data interface{}
	r.ToJSON(&data)
	return data, err
}

func (hc *HttpClient) Put(url string, paramData map[string]interface{}, headerParam map[string]string) (interface{}, error) {
	param, header := hc.handle(paramData, headerParam)
	r, err := req.Put(url, param, header)
	if err != nil {
		setting.Logger.Infof("http:request:"+url+"===error", err)
	}
	var data interface{}
	r.ToJSON(&data)
	return data, err
}

func (hc *HttpClient) Delete(url string, paramData map[string]interface{}, headerParam map[string]string) (interface{}, error) {
	param, header := hc.handle(paramData, headerParam)
	r, err := req.Delete(url, param, header)
	if err != nil {
		setting.Logger.Infof("http:request:"+url+"===error", err)
	}
	var data interface{}
	r.ToJSON(&data)
	return data, err
}

func (hc *HttpClient) Get(url string, paramData map[string]interface{}, headerParam map[string]string) (interface{}, error) {
	param, header := hc.handle(paramData, headerParam)
	r, err := req.Get(url, param, header)
	if err != nil {
		setting.Logger.Infof("http:request:"+url+"===error", err)
	}
	var data interface{}
	r.ToJSON(&data)
	return data, err
}

func (hc *HttpClient) PostBodyJson(url string, body interface{}) (interface{}, error) {
	r, err := req.Post(url, req.BodyJSON(&body))
	if err != nil {
		setting.Logger.Infof("http:request:"+url+"===error", err)
	}
	var data interface{}
	r.ToJSON(&data)
	return data, err
}

func (hc *HttpClient) PostBodyXml(url string, body interface{}) (interface{}, error) {
	r, err := req.Post(url, req.BodyXML(&body))
	if err != nil {
		setting.Logger.Infof("http:request:"+url+"===error", err)
	}
	var data interface{}
	r.ToJSON(&data)
	return data, err
}

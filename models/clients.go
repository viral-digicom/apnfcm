package models

import (
	"fmt"
	"log"
	"os"
	"bytes"
	"net/http"
	"encoding/json"
	"golang.org/x/net/http2"
)
const (
	androidDevHost      = "https://fcm.googleapis.com/fcm/send"
	devAppleHost        = "https://api.development.push.apple.com:443"
	productionAppleHost = "https://api.push.apple.com:443"
)

type APNSClient struct {
	URL        string
	HTTPClient *http.Client
	Logger     *log.Logger
}

func NewClient(isSanbox bool) (*APNSClient, error) {
	var urlString string
	if isSanbox {
		urlString = devAppleHost
	} else {
		urlString = productionAppleHost
	}
	tr := &http.Transport{}
	err := http2.ConfigureTransport(tr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	client := &http.Client{
		Transport: tr,
	}
	logger := log.New(os.Stdout, "ppush ", log.LstdFlags)
	apnsClient := &APNSClient{
		URL:        urlString,
		HTTPClient: client,
		Logger:     logger,
	}
	return apnsClient, nil
}

func NewAndroidClient() (*APNSClient, error) {
	var urlString string
	urlString = androidDevHost
	tr := &http.Transport{}
	err := http2.ConfigureTransport(tr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	client := &http.Client{
		Transport: tr,
	}
	logger := log.New(os.Stdout, "ppush ", log.LstdFlags)
	apnsClient := &APNSClient{
		URL:        urlString,
		HTTPClient: client,
		Logger:     logger,
	}
	return apnsClient, nil
}

func NewHeader(authorization string, topic string) Header {
	return Header{
		Authorization: authorization,
		Topic:         topic,
	}
}

func NewAndroidHeader(authorization string) AndroidHeader {
	return AndroidHeader{
		Authorization: authorization,
	}
}

func (c *APNSClient) APNsRequest(token string, header Header, payload IOSAPS) (*http.Request, error) {
	URL := fmt.Sprintf("%s/3/device/%s", c.URL, token)
	b, err := json.Marshal(payload.Map())
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", URL, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header = header.Map()
	return req, nil
}

func (c *APNSClient) AndroidsRequest(header AndroidHeader, payload AndroidAPN) (*http.Request, error) {
	URL := c.URL
	b, err := json.Marshal(payload.Map())
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", URL, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	/*hdr := make(map[string]string, 6)
	hdr["authorization"] = header.Authorization
	hdr["Content-Type"] = "application/json"*/
	req.Header = header.Map()
	/*for k := range hdr {
		req.Header.Add(k, string(hdr[k]))
		fmt.Printf("key[%s] value[%s]\n", k, hdr[k])
	}*/
	return req, nil
}

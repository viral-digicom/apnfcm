package apnfcm

import (
	"apnfcm/jwt"
	"apnfcm/models"
	"encoding/json"
	"fmt"
	"github.com/kataras/go-errors"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Token   string
	Message string
	IsError bool
}
type AndroidResponse struct {
	Message string
	Tokens  []string
	IsError bool
}

type iOS struct {
	privateKeyPath string
	teamId         string
	keyId          string
	topic          string
	jwtToken       string
}

type android struct {
	FCMToken string
}

var iosModal iOS
var androidModal android
var deviceType int = 0

func InitIosAPN(privateKeyPath string, teamId string, keyId string, topic string) error {
	iosModal.teamId = teamId
	iosModal.privateKeyPath = privateKeyPath
	iosModal.keyId = keyId
	iosModal.topic = topic
	jwtToken, err := jwt.CreateJWT(iosModal.privateKeyPath, iosModal.keyId, iosModal.teamId)
	if err != nil {
		return err
	}
	iosModal.jwtToken = jwtToken
	deviceType = 1
	return nil
}

func InitAndroidAPN(FCMKey string) {
	androidModal.FCMToken = FCMKey
	deviceType = 2
}

func SendAndroid(androidNotificationModal models.AndroidAPN) ([]AndroidResponse, error) {
	if deviceType != 2 {
		return nil, errors.New("Please initialize Android Package");
	}
	var req *http.Request
	var res *http.Response
	apnClient, err := models.NewAndroidClient()
	if err != nil {
		return nil, err
	}
	error := make(chan AndroidResponse)
	header := models.NewAndroidHeader("key=" + androidModal.FCMToken)
	resp := make([]AndroidResponse, 0)
	req, err = apnClient.AndroidsRequest(header, androidNotificationModal)
	go androidExecuteRequest(apnClient, req, res, androidNotificationModal.RegistrationIds, error)
	for customResponse := range error {
		resp = append(resp, customResponse)
		close(error)
	}
	return resp, nil
}

func SendIOS(deviceIds []string, iosNotificationModal models.IOSAPS) ([]Response, error) {
	if deviceType != 1 {
		return nil, errors.New("Please initialize ios private keys");
	}
	var req *http.Request
	var res *http.Response
	apnClient, err := models.NewClient(true)
	if err != nil {
		return nil, err
	}
	error := make(chan Response)
	header := models.NewHeader(iosModal.jwtToken, iosModal.topic)
	resp := make([]Response, 0)
	for _, token := range deviceIds {
		req, err = apnClient.APNsRequest(token, header, iosNotificationModal)
		go executeRequest(apnClient, req, res, token, error)
	}
	var dd = 0
	for customResponse := range error {
		dd += 1
		resp = append(resp, customResponse)
		if dd == len(deviceIds) {
			close(error)
		}
	}
	return resp, nil
}

func executeRequest(client *models.APNSClient, req *http.Request, res *http.Response, token string, error chan Response) {
	res, err := client.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
		error <- Response{err.Error(), token, true}
		return
	}
	fmt.Println(res.Status)
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if err != nil {
		fmt.Println(err)
		var r string
		if err := json.Unmarshal(body, &r); err != nil {
			fmt.Println(err)
			error <- Response{err.Error(), token, true}
			return
		}
		fmt.Println(r)
		error <- Response{err.Error(), token, true}
		return
	}
	error <- Response{"Success", token, false}
}
func androidExecuteRequest(client *models.APNSClient, req *http.Request, res *http.Response, token []string, error chan AndroidResponse) {
	res, err := client.HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
		error <- AndroidResponse{err.Error(), token, true}
		return
	}
	fmt.Println(res.Status)
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	if err != nil {
		fmt.Println(err)
		var r string
		if err := json.Unmarshal(body, &r); err != nil {
			fmt.Println(err)
			error <- AndroidResponse{err.Error(), token, true}
			return
		}
		fmt.Println(r)
		error <- AndroidResponse{err.Error(), token, true}
		return
	}
	error <- AndroidResponse{string(body), token, false}
}

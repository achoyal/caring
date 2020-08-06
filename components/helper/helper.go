package helper

import (
	"errors"
	"log"
	"net/http"
	"time"
)

type ServiceParams struct {
	URI     string
	Method  string
	Token   string
	ReqBody []byte
	Timeout time.Duration
}

//CallService
func CallService(serviceParams ServiceParams) (response *http.Response, err error) {
	switch serviceParams.Method {
	case http.MethodGet:
		{
			return GetRequest(serviceParams.URI, serviceParams.Token)
		}
	case http.MethodPost:
		{
			//return PostJson(serviceParams)
		}
	default:
		return nil, errors.New("Not a valid Method")
	}
	return
}

func GetRequest(url string, authHeader string) (response *http.Response, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return
	}
	req.Header.Add("username", authHeader)
	client := http.Client{
		//Timeout: time.Duration(config.NotificationService.Timeout * time.Second),
	}
	log.Println("req", req)
	response, err = client.Do(req)
	return
}

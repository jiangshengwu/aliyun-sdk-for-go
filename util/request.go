package util

import (
	"io/ioutil"
	"log"
	"net/http"
)

type AliyunRequest struct {
	Url string
}

func (request *AliyunRequest) DoGetRequest() (string, error) {
	resp, err := http.Get(request.Url)
	if err != nil {
		// handle error
		log.Fatal(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		log.Fatal(err)
		return "", err
	}
	result := string(body)
	return result, err
}

func GetQueryFromMap(params map[string]string) string {
	query := ""
	for k, v := range params {
		query += "&" + k + "=" + v
	}
	return query
}

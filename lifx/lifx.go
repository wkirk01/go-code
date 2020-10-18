package lifx

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const lifxBaseURL = "https://api.lifx.com/v1/"

// var token = "c4682413f1c4efa36095a51f580fc31d45a2f461b4649126e7061173b4984bc3"

var token string

//CreateLifxClient ...
func InitClient(appToken string) {
	token = appToken
}

// MakeRequest ...
func MakeRequest(method string, endpoint string, body io.Reader) []byte {
	client := http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, lifxBaseURL+endpoint, body)
	check(err)
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	check(err)
	return respBody
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

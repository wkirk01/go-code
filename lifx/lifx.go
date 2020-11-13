package lifx

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const lifxBaseURL = "https://api.lifx.com/v1/"

// var token = "c4682413f1c4efa36095a51f580fc31d45a2f461b4649126e7061173b4984bc3"

// Client ...
type Client struct {
	token string
}

// NewClient ...
func NewClient(apptoken string) *Client {
	return &Client{token: apptoken}
}

// MakeRequest ...
func (c *Client) MakeRequest(method string, endpoint string, body io.Reader) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, lifxBaseURL+endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

//GetLights ...
func (c *Client) GetLights() ([]Light, error) {
	respBody, err := c.MakeRequest("GET", "lights/all", nil)
	if err != nil {
		return nil, err
	}
	lights := []Light{}
	err = json.Unmarshal(respBody, &lights)
	if err != nil {
		return nil, err
	}
	return lights, nil
}

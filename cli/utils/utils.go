package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Config *ConfigClient
}

func NewClient(config *ConfigClient) *Client {
	return &Client{
		Config: config,
	}
}

func (c *Client) PingServer() (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s:%s/ping", c.Config.Server.Host, c.Config.Server.Port))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}
	return string(body), nil
}

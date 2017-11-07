package main

import (
	"encoding/json"
	"fmt"
	"net/http"

"github.com/gopherjs/gopherjs/js"
)

type Config struct {
	GrpcHost string `json:"grpc_host"`
}

func LoadConfig() (*Config, error) {
	scheme := js.Global.Get("location").Get("protocol").String()
	hostname := js.Global.Get("location").Get("hostname").String()


	response, err := http.Get(fmt.Sprintf("%s//%s/config.json", scheme, hostname))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	config := new(Config)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

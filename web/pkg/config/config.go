package config

import (
	"github.com/cathalgarvey/fmtless"
	"github.com/cathalgarvey/fmtless/encoding/json"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/xhr"
)

type Config struct {
	GrpcHost string `json:"grpc_host"`
}

func Load() (*Config, error) {
	scheme := js.Global.Get("location").Get("protocol").String()
	hostname := js.Global.Get("location").Get("hostname").String()

	data, err := xhr.Send("GET", fmt.Sprintf("%s//%s/config.json", scheme, hostname), nil)
	if err != nil {
		return nil, err
	}

	config := new(Config)
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

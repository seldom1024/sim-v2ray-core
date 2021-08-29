package io

import (
	"encoding/json"
	core "github.com/seldom1024/sim-v2ray-core"
)

type JsonVUser struct {
	id    string `json:"id"`
	email string `json:"email"`
}

type JsonVConfig struct {
	Port     uint8       `json:"port"`
	Clients  []JsonVUser `json:"users"`
	Protocol string      `json:"protocol"`
}

type JsonVConfigUnmarshaller struct {
}

func (*JsonVConfigUnmarshaller) Unmarshall(data []byte) (*core.VConfig, error) {
	var jsonConfig JsonVConfig
	err := json.Unmarshal(data, &jsonConfig)
	if err != nil {
		return nil, err
	}
	var vconfig = new(core.VConfig)
	return vconfig, nil
}

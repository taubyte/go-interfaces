package common

import (
	"context"
)

type CommonConfig struct {
	Disabled bool
	Port     int
	Root     string
}

type ServiceConfig struct {
	CommonConfig
	Ctx    context.Context
	Others map[string]int
}

type SimpleConfig struct {
	CommonConfig
	Clients map[string]ClientConfig
}

func (c *ServiceConfig) Clone() *ServiceConfig {
	cclone := new(ServiceConfig)
	cclone.CommonConfig = c.CommonConfig
	cclone.Ctx = c.Ctx
	cclone.Others = make(map[string]int)
	for k, v := range c.Others {
		cclone.Others[k] = v
	}
	return cclone
}

type ClientConfig struct {
	CommonConfig
}

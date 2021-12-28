package http

import (
	cHttp "github.com/kldx/common/http"
	cStructs "github.com/kldx/common/structs"
	cUtils "github.com/kldx/common/utils"
	"net/http"
	"sync"
	"time"
)

var (
	openapiClientOnce sync.Once
	openapiClient     *cHttp.HttpClient

	openapiClientConf = &cStructs.HttpConfig{
		Url:                 cUtils.GetOpenapiUrl(),
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     10 * time.Second,
	}
)

func GetOpenapiClient() *cHttp.HttpClient {
	openapiClientOnce.Do(func() {
		openapiClient = &cHttp.HttpClient{
			Client: http.Client{
				Transport: &http.Transport{
					MaxIdleConns:        openapiClientConf.MaxIdleConns,
					MaxIdleConnsPerHost: openapiClientConf.MaxIdleConnsPerHost,
					IdleConnTimeout:     openapiClientConf.IdleConnTimeout,
				},
			},
			Url: openapiClientConf.Url,
		}
	})
	return openapiClient
}

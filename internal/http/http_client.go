package http

import (
	"net/http"
	"sync"
	"time"
)

var (
	httpClientInstance *http.Client
	once               sync.Once
)

func GetHttpClient() *http.Client {

	//timeOut := config.GetConfig().HTTP.Client.Timeout

	once.Do(func() {
		httpClientInstance = &http.Client{
			Timeout: 10 * time.Second,
		}
	})
	return httpClientInstance
}

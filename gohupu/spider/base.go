package spider

import (
	"github.com/go-resty/resty/v2"
	"github.com/wudizhangzhi/HupuApp/gohupu/logger"
	"net/http"
)

var SpiderClient *resty.Client

func New() *resty.Client {
	SpiderClient = resty.New().
		SetPreRequestHook(func(c *resty.Client, r *http.Request) error {
			// fmt.Printf("Request: %+v\n", r.URL)
			logger.Debug.Printf("Request: %+v\n", r.URL)
			return nil
		})
	return SpiderClient
}

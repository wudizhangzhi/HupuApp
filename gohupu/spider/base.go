package spider

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

var SpiderClient *resty.Client

func New() *resty.Client {
	return resty.New().
		SetPreRequestHook(func(c *resty.Client, r *http.Request) error {
			fmt.Printf("Request: %+v\n", r.URL)
			return nil
		})
}

package api

import (
	"fmt"
	"math/rand"
	"net/http"
	net_url "net/url"
	"strings"
	"time"

	"github.com/wudizhangzhi/HupuApp"
)

var HupuHttpobj HupuHttp

type HupuHttp struct {
	Client *http.Client
	// Method  string
	// Params  map[string]string
	Headers   map[string]string
	IMEI      string
	AndroidId string // 也是clientid
}

func init() {
	rand.Seed(time.Now().Unix())
	agent := HupuApp.ANDROID_USER_AGENT[rand.Intn(len(HupuApp.ANDROID_USER_AGENT))]
	HupuHttpobj = HupuHttp{
		Client: &http.Client{},
		Headers: map[string]string{
			"User-Agent": agent + " kanqiu/" + HupuApp.API_VERSION + ".13305/7214 isp/-1 network/-1",
		},
		IMEI:      HupuApp.GetRandomImei(0, ""),
		AndroidId: HupuApp.GetAndroidId(),
	}
	fmt.Printf("初始化: %+v\n", HupuHttpobj)
}

func (hh *HupuHttp) Request(method string, url string, headers map[string]string, params map[string]string) (*http.Response, error) {
	fmt.Printf("访问接口: [%s] %s\n", method, url)
	sign := HupuApp.GetSortParam(params)
	params["sign"] = sign
	var req *http.Request
	var err error
	switch method {
	case "POST":
		// data := make(map[string][]string)
		data := net_url.Values{}
		for k, v := range params {
			data[k] = []string{v}
		}
		req, err = http.NewRequest(method, url, strings.NewReader(data.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		if err != nil {
			return nil, err
		}
	case "GET":
		// format url param
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
		fmt.Printf("参数: %s\n", req.URL.RawQuery)
	}

	for k, v := range hh.Headers {
		req.Header.Set(k, v)
	}
	fmt.Printf("headers: %+v\n", req.Header)
	return hh.Client.Do(req)
}

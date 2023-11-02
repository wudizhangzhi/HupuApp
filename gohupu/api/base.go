package api

import (
	"math/rand"
	"net/http"
	net_url "net/url"
	"strings"

	"github.com/wudizhangzhi/HupuApp"
	"github.com/wudizhangzhi/HupuApp/gohupu/logger"
)

var HttpSession Session

type Session struct {
	Client    *http.Client
	Headers   map[string]string
	IMEI      string
	AndroidId string // 也是clientid
}

func init() {
	// 初始化虎扑专用http连接
	HttpSession = *NewSession()
	logger.Info.Printf("初始化: %+v\n", HttpSession)
}

func NewSession() *Session {
	agent := HupuApp.ANDROID_USER_AGENT[rand.Intn(len(HupuApp.ANDROID_USER_AGENT))]
	return &Session{
		Client: &http.Client{},
		Headers: map[string]string{
			"User-Agent": agent + " kanqiu/" + HupuApp.API_VERSION + ".13305/7214 isp/-1 network/-1",
		},
		IMEI:      HupuApp.GetRandomImei(0, ""),
		AndroidId: HupuApp.GetAndroidId(),
	}
}

func (s *Session) Request(method string, url string, headers map[string]string, params map[string]string) (*http.Response, error) {
	logger.Info.Printf("访问接口: [%s] %s\n", method, url)
	sign := HupuApp.GetSortParam(params)
	params["sign"] = sign
	var req *http.Request
	var err error
	switch method {
	case "POST":
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
		logger.Info.Printf("参数: %s\n", req.URL.RawQuery)
	}

	for k, v := range s.Headers {
		req.Header.Set(k, v)
	}
	logger.Info.Printf("headers: %+v\n", req.Header)
	return s.Client.Do(req)
}

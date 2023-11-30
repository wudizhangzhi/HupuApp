package api

import (
	"math/rand"

	"github.com/go-resty/resty/v2"
	"github.com/wudizhangzhi/HupuApp"
	"github.com/wudizhangzhi/HupuApp/gohupu/logger"
)

var HttpSession Session

type Session struct {
	// Client    *http.Client
	Client    *resty.Client
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
		Client: resty.New(),
		Headers: map[string]string{
			"User-Agent": agent + " kanqiu/" + HupuApp.API_VERSION + ".13305/7214 isp/-1 network/-1",
		},
		IMEI:      HupuApp.GetRandomImei(0, ""),
		AndroidId: HupuApp.GetAndroidId(),
	}
}

func (s *Session) Request(method string, url string, headers map[string]string, params map[string]string) (*resty.Response, error) {
	logger.Info.Printf("访问接口: [%s] url:%s headers: %+v, 参数: %v \n", method, url, headers, params)
	sign := HupuApp.GetSortParam(params)
	params["sign"] = sign
	var err error
	var resp *resty.Response
	switch method {
	case "POST":
		resp, err = s.Client.R().
			SetBody(params).
			SetHeaders(s.Headers).
			SetHeader("Content-Type", "application/x-www-form-urlencoded").
			Post(url)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "GET":
		resp, err = s.Client.R().
			SetHeaders(s.Headers).
			SetQueryParams(params).
			Get(url)
		if err != nil {
			return nil, err
		}
	}
	return resp, err
}

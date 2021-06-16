package api

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"net/http"
	net_url "net/url"
	"sort"
	"strings"
	"time"

	"github.com/wudizhangzhi/HupuApp"
)

var HupuHttpobj HupuHttp

type HupuHttp struct {
	Client *http.Client
	// Method  string
	// Params  map[string]string
	Headers map[string]string
}

// def getSortParam(**kwargs):
//     result = ''
//     kwargs_sorted = sorted(kwargs)
//     for key in kwargs_sorted:
//         if len(result) > 0:
//             result += '&'
//         result += '='.join((key, str(kwargs.get(key))))
//     result += HUPU_SALT
//     return md5(result.encode('utf8')).hexdigest()

func init() {
	rand.Seed(time.Now().Unix())
	agent := HupuApp.ANDROID_USER_AGENT[rand.Intn(len(HupuApp.ANDROID_USER_AGENT))]
	HupuHttpobj = HupuHttp{
		Client: &http.Client{},
		Headers: map[string]string{
			"User-Agent": agent,
		},
	}
}

func getSortParam(params map[string]string) string {
	result := ""
	keys := make([]string, 0)
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		if len(result) > 0 {
			result += "&"
		}
		result += strings.Join([]string{key, params[key]}, "=")
	}
	result += HupuApp.HUPU_SALT
	h := md5.New()
	h.Write([]byte(result))
	return hex.EncodeToString(h.Sum(nil))
}

func (hh *HupuHttp) Request(method string, url string, headers map[string]string, params map[string]string) (*http.Response, error) {
	sign := getSortParam(params)
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
		if err != nil {
			return nil, err
		}

		// return hh.Client.PostForm(url, data)
	case "GET":
		// format url param
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
		// return hh.Client.Do(req)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range hh.Headers {
		req.Header.Set(k, v)
	}
	return hh.Client.Do(req)
}

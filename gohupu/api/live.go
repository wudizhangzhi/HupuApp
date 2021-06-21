package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/wudizhangzhi/HupuApp"
	"github.com/wudizhangzhi/HupuApp/gohupu/message"
)

var (
	Domain string
)

func init() {
	// 初始化数值
	addresses := GetIpAddress()
	Domain = HupuApp.RandomChoice(addresses)
}

// url='https://games.mobileapi.hupu.com/1/{}/status/init'.format(self.api_version),
//             params={
//                 'dv': '5.7.79',
//                 'crt': int(time.time() * 1000),
//                 'tag': 'nba',  # 默认nba
//                 'night': 0,
//                 'channel': 'myapp',
//                 'client': self.client,
//                 'time_zone': 'Asia/Shanghai',
//                 'android_id': self.android_id,
//             },
func APIStatusInit() (*http.Response, error) {
	params := map[string]string{
		"div":        "5.7.79",
		"crt":        fmt.Sprint(HupuApp.GetTimestamp()),
		"tag":        "nba",
		"night":      "0",
		"channel":    "myapp",
		"client":     HupuHttpobj.IMEI,
		"time_zone":  "Asia/Shanghai",
		"android_id": HupuHttpobj.AndroidId,
	}
	return HupuHttpobj.Request("GET", HupuApp.API_STATUS_INIT, nil, params)
}

// 获取接口ip地址
func GetIpAddress() []string {
	resp, err := APIStatusInit()
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	addressJson := jsoniter.Get(respBody, "result", "redirector").ToString()
	var address []string
	json.Unmarshal([]byte(addressJson), &address)
	return address
}

// 获取比赛直播信息
func APIGetPlayByPlay(gid int) (*http.Response, error) {
	params := map[string]string{
		"gid":        fmt.Sprint(gid),
		"lid":        "1",
		"roomid":     "-1",
		"entrance":   "-1",
		"tag":        "nba",
		"channel":    "myapp",
		"night":      "0",
		"crt":        fmt.Sprint(HupuApp.GetTimestamp()),
		"client":     HupuHttpobj.IMEI,
		"time_zone":  "Asia/Shanghai",
		"android_id": HupuHttpobj.AndroidId,
	}
	return HupuHttpobj.Request("GET", HupuApp.API_GET_PLAY_BY_PLAY, nil, params)
}

type GameType string

const (
	NBA GameType = "nba"
	CBA GameType = "cba"
)

// 获取比赛列表
func APIGetGames(gametype GameType) (*http.Response, error) {
	params := map[string]string{
		"night":      "0",
		"channel":    "myapp",
		"crt":        fmt.Sprint(HupuApp.GetTimestamp()),
		"client":     HupuHttpobj.IMEI,
		"time_zone":  "Asia/Shanghai",
		"android_id": HupuHttpobj.AndroidId,
	}
	return HupuHttpobj.Request("GET", fmt.Sprintf(HupuApp.API_GET_GAMES, gametype), nil, params)
}

func GetDayGames(gametype GameType) ([]message.DayGame, error) {
	results := []message.DayGame{}
	resp, err := APIGetGames(gametype)
	if err != nil {
		return results, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return results, err
	}
	// fmt.Println("response Body : ", string(respBody))
	msgGame := message.MsgGame{}
	if err := json.Unmarshal(respBody, &msgGame); err != nil {
		return results, err
	}
	return msgGame.Result.DayGames, nil
}

// 获取今天的比赛
func GetGameToday(gametype GameType) ([]message.Game, error) {
	result := make([]message.Game, 0)
	dayGames, err := GetDayGames(gametype)
	if err != nil {
		return result, nil
	}
	today := time.Now().Format("20060102")
	for _, dayGame := range dayGames {
		if dayGame.Day == today {
			result = append(result, dayGame.Games...)
		}
	}
	return result, nil
}

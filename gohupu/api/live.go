package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tidwall/gjson"

	"github.com/wudizhangzhi/HupuApp"
	"github.com/wudizhangzhi/HupuApp/gohupu/logger"
	"github.com/wudizhangzhi/HupuApp/gohupu/message"
)

var (
	Domain string
)

func init() {
	// 初始化数值
	addresses := GetIpAddress()
	logger.Info.Printf("获取所有地址: %s\n", addresses)
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
	addressJson := gjson.GetBytes(respBody, "result.redirector").String()

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

// 获取今天的比赛
func GetMatchesToday(gametype GameType) ([]message.Match, error) {
	result := make([]message.Match, 0)
	schedule, err := GetScheduleList(gametype)
	if err != nil {
		return result, err
	}
	today := time.Now().Format("20060102")
	for _, dayGame := range schedule.GameList {
		if dayGame.Day == today {
			result = append(result, dayGame.MatchList...)
		}
	}
	return result, nil
}

func GetMatchesFromDate(gametype GameType, dates ...string) ([]message.Match, error) {
	matches := make([]message.Match, 0)
	schedule, err := GetScheduleList(gametype)
	if err != nil {
		return matches, nil
	}
	var date string
	if len(dates) == 0 {
		date = time.Now().Format("20060102")
	} else {
		date = dates[0]
	}

	for _, game := range schedule.GameList {
		logger.Info.Printf("对比日期: %s - %s", date, game.Day)
		if game.Day == date {
			matches = append(matches, game.MatchList...)
			break
		}
	}
	return matches, nil
}

func APIGetScheduleList(gametype GameType) (*http.Response, error) {
	params := map[string]string{
		"competitionTag": string(gametype),
		"night":          "0",
		"V":              "7.5.59.01043",
		"channel":        "hupuupdate",
		"crt":            fmt.Sprint(HupuApp.GetTimestamp()),
		"_imei":          HupuHttpobj.IMEI,
		"time_zone":      "Asia/Shanghai",
		"android_id":     HupuHttpobj.AndroidId,
		// "client":     HupuHttpobj.IMEI,
	}
	return HupuHttpobj.Request("GET", HupuApp.API_SCHEDULE_LIST, nil, params)
}

func GetScheduleList(gametype GameType) (message.GameSchedule, error) {
	gameSchdule := message.GameSchedule{}
	resp, err := APIGetScheduleList(gametype)
	if err != nil {
		return gameSchdule, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return gameSchdule, err
	}

	byteResult, _ := json.Marshal(gjson.GetBytes(respBody, "result").Value())
	json.Unmarshal(byteResult, &gameSchdule)

	// logger.Info.Printf("ScheduleList返回: %v", gameSchdule)
	return gameSchdule, nil
}

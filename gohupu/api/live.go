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

type GameType string

const (
	NBA GameType = "nba"
	CBA GameType = "cba"
)

func init() {
	// 初始化数值
	addresses := GetIpAddress()
	logger.Info.Printf("获取所有地址: %s\n", addresses)
	Domain = HupuApp.RandomChoice(addresses)
}

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

// 获取直播的key
func APIQueryLiveActivityKey(matchId string) (*http.Response, error) {
	params := map[string]string{
		"competitionType": "basketball",
		"matchId":         matchId,
		"channel":         "hupuupdate",
		"night":           "0",
		"crt":             fmt.Sprint(HupuApp.GetTimestamp()),
		"_imei":           HupuHttpobj.IMEI,
		"time_zone":       "Asia/Shanghai",
		"android_id":      HupuHttpobj.AndroidId,
	}
	return HupuHttpobj.Request("GET", HupuApp.API_LIVE_QUERY_LIVE_ACTIVITY_KEY, nil, params)
}

// 获取直播内容
func APIQueryLiveTextList(matchId string, liveActivityKeyStr string, commentId string) (*http.Response, error) {
	params := map[string]string{
		"competitionType":    "basketball",
		"matchId":            matchId,
		"liveActivityKeyStr": liveActivityKeyStr,
		"channel":            "hupuupdate",
		"night":              "0",
		"crt":                fmt.Sprint(HupuApp.GetTimestamp()),
		"_imei":              HupuHttpobj.IMEI,
		"time_zone":          "Asia/Shanghai",
		"android_id":         HupuHttpobj.AndroidId,
	}
	if commentId != "" {
		params["commentId"] = commentId
	}
	return HupuHttpobj.Request("GET", HupuApp.API_LIVE_QUERY_LIVE_TEXT_LIST, nil, params)
}

// 获取赛程
func APIGetScheduleList(gametype GameType, coursors ...string) (*http.Response, error) {
	params := map[string]string{
		"competitionTag": fmt.Sprint(gametype),
		"night":          "0",
		"V":              "7.5.59.01043",
		"channel":        "hupuupdate",
		"crt":            fmt.Sprint(HupuApp.GetTimestamp()),
		"_imei":          HupuHttpobj.IMEI,
		"time_zone":      "Asia/Shanghai",
		"android_id":     HupuHttpobj.AndroidId,
	}
	if len(coursors) > 0 {
		params["coursor"] = coursors[0]
	}
	return HupuHttpobj.Request("GET", HupuApp.API_SCHEDULE_LIST, nil, params)
}

// 获取单个比赛信息
func APISingleMatch(matchId string) (*http.Response, error) {
	params := map[string]string{
		"matchId":    matchId,
		"night":      "0",
		"V":          "7.5.59.01043",
		"channel":    "hupuupdate",
		"crt":        fmt.Sprint(HupuApp.GetTimestamp()),
		"_imei":      HupuHttpobj.IMEI,
		"time_zone":  "Asia/Shanghai",
		"android_id": HupuHttpobj.AndroidId,
	}
	return HupuHttpobj.Request("GET", HupuApp.API_SCHEDULE_LIST, nil, params)
}

// 获取某天的比赛
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

func GetAnyMatches(gameType GameType) ([]message.Match, error) {
	matches := make([]message.Match, 0)
	schedule, err := GetScheduleList(gameType)
	if err != nil {
		return matches, nil
	}
	for _, game := range schedule.GameList {
		if len(matches) > 5 {
			break
		}
		matches = append(matches, game.MatchList...)
	}
	return matches, nil
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
	return gameSchdule, nil
}

func QueryLiveTextList(matchId string, liveActivityKeyStr string, commentId string) ([]message.MatchTextMsg, error) {
	matchTextMsgs := []message.MatchTextMsg{}

	resp, err := APIQueryLiveTextList(matchId, liveActivityKeyStr, commentId)
	if err != nil {
		return matchTextMsgs, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return matchTextMsgs, err
	}
	for _, msg := range gjson.GetBytes(respBody, "result").Array() {
		matchTextMsg := message.MatchTextMsg{}
		byteResult, _ := json.Marshal(msg.Value())
		json.Unmarshal(byteResult, &matchTextMsg)
		matchTextMsgs = append(matchTextMsgs, matchTextMsg)
		logger.Info.Printf("比赛消息: %v", msg.Value())
	}
	return matchTextMsgs, nil
}

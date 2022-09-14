package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"time"

	"github.com/tidwall/gjson"

	"github.com/wudizhangzhi/HupuApp"
	"github.com/wudizhangzhi/HupuApp/gohupu/logger"
	"github.com/wudizhangzhi/HupuApp/gohupu/message"
)

type GameType string

const (
	NBA GameType = "nba"
	CBA GameType = "cba"
)

func init() {
	// 初始化数值
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
		// competitionType	basketball
		// matchId	980817952171360256
		// clientId	133092838
		// deviceId	Btv88Xr1R0lJNZTOalfXGk6/BR3oYpZ71Gu/4xLaRgU8PNeg+GmrInk2BSXy1H9mCGcIHFr5FQb9eenF9Rh5VcQ==
		// _ssid	IktLUHZteEJwMHEi
		// _imei	868364056062517
		// time_zone	Asia/Shanghai
		// client	69bf46d0f4e96e84
		// night	0
		// crt	1662706885934
		// channel	hupuupdate
		// android_id	69bf46d0f4e96e84
		// sign	25af1c868310a9ecec0645e6b141eabc
	}
	return HupuHttpobj.Request("GET", HupuApp.API_LIVE_QUERY_LIVE_ACTIVITY_KEY, nil, params)
}

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
		// matchId	980817952171360256
		// liveActivityKeyStr	32476308:10922:261542
		// clientId	133092838
		// deviceId	Btv88Xr1R0lJNZTOalfXGk6/BR3oYpZ71Gu/4xLaRgU8PNeg+GmrInk2BSXy1H9mCGcIHFr5FQb9eenF9Rh5VcQ==
		// _ssid	IktLUHZteEJwMHEi
		// _imei	868364056062517
		// time_zone	Asia/Shanghai
		// client	69bf46d0f4e96e84
		// night	0
		// crt	1662706885934
		// channel	hupuupdate
		// android_id	69bf46d0f4e96e84
		// sign	737852131bf836fd2ff61780507103af
	}
	if commentId != "" {
		params["commentId"] = commentId
	}
	return HupuHttpobj.Request("GET", HupuApp.API_LIVE_QUERY_LIVE_TEXT_LIST, nil, params)
}

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
		// "client":     HupuHttpobj.IMEI,
		// TODO 测试
		"coursor": "20220718",
	}
	if len(coursors) > 0 {
		params["coursor"] = coursors[0]
	}
	return HupuHttpobj.Request("GET", HupuApp.API_SCHEDULE_LIST, nil, params)
}

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
		// "client":     HupuHttpobj.IMEI,
	}
	return HupuHttpobj.Request("GET", HupuApp.API_SINGLE_MATCH, nil, params)
}

func GetSingleMatch(matchId string) (message.Match, error) {
	match := message.Match{}

	resp, err := APISingleMatch(matchId)
	if err != nil {
		return match, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return match, err
	}
	result := gjson.GetBytes(respBody, "result").Value()
	byteResult, _ := json.Marshal(result)
	json.Unmarshal(byteResult, &match)
	logger.Info.Printf("比赛状态: %+v", result)
	return match, nil
}

// 获取比赛
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

func GetAnyMatches(gameType GameType, count int, reverse bool) ([]message.Match, error) {
	today := time.Now()
	matches := make([]message.Match, 0)
	schedule, err := GetScheduleList(gameType)
	if err != nil {
		return matches, nil
	}
	for _, game := range schedule.GameList {
		t, _ := time.Parse("20060102", game.Day)
		if t.Unix() <= today.Unix() {
			for _, match := range game.MatchList {

				matches = append(matches, match)
			}
		}
	}
	sort.Slice(matches, func(i, j int) bool {
		if reverse {
			return matches[i].ChinaStartTime < matches[j].ChinaStartTime
		} else {
			return matches[i].ChinaStartTime > matches[j].ChinaStartTime
		}
	})
	if count < len(matches)-1 {
		matches = matches[:count]
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

	// DEBUG
	total := 0
	for _, game := range gameSchdule.GameList {
		total += len(game.MatchList)
	}
	logger.Info.Printf("ScheduleList返回: %d 天 %d 场比赛", len(gameSchdule.GameList), total)
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

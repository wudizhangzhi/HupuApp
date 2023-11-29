package api

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/wudizhangzhi/HupuApp"
)

type GameType string

const (
	NBA GameType = "nba"
	CBA GameType = "cba"
)

func (g GameType) String() string {
	switch g {
	case NBA:
		return "nba"
	case CBA:
		return "cba"
	default:
		return "nba"
	}
}

func init() {

}

// 状态初始化，获取基本信息接口
func GetInitInfo() (*resty.Response, error) {
	params := map[string]string{
		"div":        "5.7.79",
		"crt":        fmt.Sprint(HupuApp.GetTimestamp()),
		"tag":        "nba",
		"night":      "0",
		"channel":    "myapp",
		"client":     HttpSession.IMEI,
		"time_zone":  "Asia/Shanghai",
		"android_id": HttpSession.AndroidId,
	}
	return HttpSession.Request("GET", HupuApp.API_STATUS_INIT, nil, params)
}

// 获取比赛直播信息
func GetPlayByPlay(gid int) (*resty.Response, error) {
	params := map[string]string{
		"gid":        fmt.Sprint(gid),
		"lid":        "1",
		"roomid":     "-1",
		"entrance":   "-1",
		"tag":        "nba",
		"channel":    "myapp",
		"night":      "0",
		"crt":        fmt.Sprint(HupuApp.GetTimestamp()),
		"client":     HttpSession.IMEI,
		"time_zone":  "Asia/Shanghai",
		"android_id": HttpSession.AndroidId,
	}
	return HttpSession.Request("GET", HupuApp.API_GET_PLAY_BY_PLAY, nil, params)
}

// 获取比赛直播信息
func GetLiveActivityKey(matchId string) (*resty.Response, error) {
	params := map[string]string{
		"competitionType": "basketball",
		"matchId":         matchId,
		"channel":         "hupuupdate",
		"night":           "0",
		"crt":             fmt.Sprint(HupuApp.GetTimestamp()),
		"_imei":           HttpSession.IMEI,
		"time_zone":       "Asia/Shanghai",
		"android_id":      HttpSession.AndroidId,
	}
	return HttpSession.Request("GET", HupuApp.API_LIVE_QUERY_LIVE_ACTIVITY_KEY, nil, params)
}

// 获取直播内容接口
func GetLiveMsgList(matchId string, liveActivityKeyStr string, commentId string) (*resty.Response, error) {
	params := map[string]string{
		"competitionType":    "basketball",
		"matchId":            matchId,
		"liveActivityKeyStr": liveActivityKeyStr,
		"channel":            "hupuupdate",
		"night":              "0",
		"crt":                fmt.Sprint(HupuApp.GetTimestamp()),
		"_imei":              HttpSession.IMEI,
		"time_zone":          "Asia/Shanghai",
		"android_id":         HttpSession.AndroidId,
	}
	if commentId != "" {
		params["commentId"] = commentId
	}
	return HttpSession.Request("GET", HupuApp.API_LIVE_QUERY_LIVE_TEXT_LIST, nil, params)
}

// 获取比赛日程列表
func GetScheduleList(gametype GameType, coursors ...string) (*resty.Response, error) {
	params := map[string]string{
		"competitionTag": gametype.String(),
		"night":          "0",
		"V":              "7.5.59.01043",
		"channel":        "hupuupdate",
		"crt":            fmt.Sprint(HupuApp.GetTimestamp()),
		"_imei":          HttpSession.IMEI,
		"time_zone":      "Asia/Shanghai",
		"android_id":     HttpSession.AndroidId,
	}
	if len(coursors) > 0 {
		params["coursor"] = coursors[0]
	}
	return HttpSession.Request("GET", HupuApp.API_SCHEDULE_LIST, nil, params)
}

// 根据比赛id获取比赛信息
func GetSingleMatch(matchId string) (*resty.Response, error) {
	params := map[string]string{
		"matchId":    matchId,
		"night":      "0",
		"V":          "7.5.59.01043",
		"channel":    "hupuupdate",
		"crt":        fmt.Sprint(HupuApp.GetTimestamp()),
		"_imei":      HttpSession.IMEI,
		"time_zone":  "Asia/Shanghai",
		"android_id": HttpSession.AndroidId,
	}
	return HttpSession.Request("GET", HupuApp.API_SINGLE_MATCH, nil, params)
}

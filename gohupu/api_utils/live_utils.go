package api_utils

import (
	"encoding/json"
	"sort"
	"time"

	"github.com/tidwall/gjson"

	"github.com/wudizhangzhi/HupuApp"
	"github.com/wudizhangzhi/HupuApp/gohupu/api"
	"github.com/wudizhangzhi/HupuApp/gohupu/logger"
	"github.com/wudizhangzhi/HupuApp/gohupu/message"
)

// 获取接口ip地址
func GetIpAddress() []string {
	resp, err := api.GetInitInfo()
	if err != nil {
		panic(err)
	}
	addressJson := gjson.GetBytes(resp.Body(), "result.redirector").String()

	var address []string
	json.Unmarshal([]byte(addressJson), &address)
	return address
}

// 获取比赛的activity key
func GetLiveActivityKey(matchId string) (string, error) {
	liveActivityKey := ""
	// 获取比赛的activity key
	resp, err := api.GetLiveActivityKey(matchId)
	if err != nil {
		return liveActivityKey, err
	}
	liveActivityKey = gjson.GetBytes(resp.Body(), "result.liveActivityKey").String()
	return liveActivityKey, nil
}

// 根据比赛id获取比赛信息
func GetSingleMatch(matchId string) (message.Match, error) {
	match := message.Match{}

	resp, err := api.GetSingleMatch(matchId)
	if err != nil {
		return match, err
	}
	result := gjson.GetBytes(resp.Body(), "result").Value()
	byteResult, _ := json.Marshal(result)
	json.Unmarshal(byteResult, &match)
	logger.Info.Printf("比赛状态: %+v", result)
	return match, nil
}

// 根据日期获取比赛
func GetMatchesFromDate(gametype api.GameType, dates ...string) ([]message.Match, error) {
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
		logger.Debug.Printf("对比日期: %s - %s", date, game.Day)
		if game.Day == date {
			matches = append(matches, game.MatchList...)
			break
		}
	}
	return matches, nil
}

// 获取最近的比赛
func GetAnyMatches(gameType api.GameType, count int, reverse bool) ([]message.Match, error) {
	today := time.Now()
	matches := make([]message.Match, 0)
	schedule, err := GetScheduleList(gameType)
	if err != nil {
		return matches, nil
	}
	for _, game := range schedule.GameList {
		t, _ := time.Parse("20060102", game.Day)
		if t.Unix() <= today.Unix() {
			matches = append(matches, game.MatchList...)
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

// 获取赛程
func GetScheduleList(gametype api.GameType) (message.GameSchedule, error) {
	gameSchdule := message.GameSchedule{}
	resp, err := api.GetScheduleList(gametype)
	if err != nil {
		return gameSchdule, err
	}

	byteResult, _ := json.Marshal(gjson.GetBytes(resp.Body(), "result").Value())
	json.Unmarshal(byteResult, &gameSchdule)

	total := 0
	for _, game := range gameSchdule.GameList {
		total += len(game.MatchList)
	}
	logger.Info.Printf("ScheduleList返回: %d 天 %d 场比赛", len(gameSchdule.GameList), total)
	return gameSchdule, nil
}

func GetLiveMsgList(matchId string, liveActivityKeyStr string, commentId string) ([]message.LiveMsg, error) {
	liveMsgs := []message.LiveMsg{}

	resp, err := api.GetLiveMsgList(matchId, liveActivityKeyStr, commentId)
	if err != nil {
		return liveMsgs, err
	}
	for _, msg := range gjson.GetBytes(resp.Body(), "result").Array() {
		liveMsg := message.LiveMsg{}
		byteResult, _ := json.Marshal(msg.Value())
		json.Unmarshal(byteResult, &liveMsg)
		// 特殊处理, 似乎一会儿是int，一会儿是string
		liveMsg.CommentId = msg.Get("commentId").String()
		liveMsg.PreviousCommentId = msg.Get("previousCommentId").String()
		liveMsgs = append(liveMsgs, liveMsg)
		logger.Info.Printf("比赛消息: %v", msg.Value())
		logger.Debug.Printf("比赛消息装载后: %v", liveMsg)
	}
	// sort
	sort.Slice(liveMsgs, func(i, j int) bool {
		ti, _ := HupuApp.GetTimeFromString(liveMsgs[i].Time)
		tj, _ := HupuApp.GetTimeFromString(liveMsgs[j].Time)
		return ti.Before((tj))
	})
	return liveMsgs, nil
}

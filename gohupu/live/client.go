package live

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
	"github.com/wudizhangzhi/HupuApp/gohupu/api"
	"github.com/wudizhangzhi/HupuApp/gohupu/logger"
	"github.com/wudizhangzhi/HupuApp/gohupu/message"
)

type Client struct {
	WsConn          *websocket.Conn
	Domain          string        // (外部接口获取的参数)
	Pid             int           // (内部接口获取的参数)
	liveActivityKey string        // (内部接口获取的参数)
	Match           message.Match // 比赛(外部接口获取的参数)
	HomeName        string        // 主队名字
	AwayName        string        // 客队名字
	Connected       bool          // ws中是否已连接
	LastTime        int           // (内部接口获取，用于比对直播数据时间)
	LastCommentId   string        //
	Th              *time.Ticker
	ErrorCh         chan interface{} // 错误channel
	OprCh           chan interface{} // 操作通道?
	InterruptCh     chan os.Signal   // 中断信号
}

// 获取token
func (c *Client) FetchToken() (string, error) {
	var token string
	client := &http.Client{}
	t := time.Now().Unix()
	url := "http://" + c.Domain + "/socket.io/1/"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		logger.Error.Print(err)
		return "", err
	}
	q := req.URL.Query()
	q.Add("client", api.HupuHttpobj.IMEI)
	q.Add("t", fmt.Sprint(t))
	q.Add("type", "1")
	q.Add("background", "false")
	req.URL.RawQuery = q.Encode()
	logger.Info.Printf("获取token: %s\n", req.URL.RawQuery)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	logger.Info.Printf("获取token接口返回: %s", string(respBody))
	token = strings.Split(string(respBody), ":50:60")[0]
	return token, nil
}

func (c *Client) init() {
	resp, err := api.APIQueryLiveActivityKey(c.Match.MatchId)
	if err != nil {
		logger.Error.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error.Fatal(err)
	}
	liveActivityKey := gjson.GetBytes(respBody, "result.liveActivityKey").String()
	c.liveActivityKey = liveActivityKey

	logger.Info.Printf("liveActivityKey: %s\n", liveActivityKey)
}

// 格式化直播消息
func (c *Client) HandleLiveMsg(msg message.MatchTextMsg) {
	// if len(msg.Args) > 0 && len(msg.Args[0].Result.Data) > 0 {
	// 	s := msg.Args[0].Result.Score
	// 	if s.AwayName != "" {
	// 		c.AwayName = s.AwayName
	// 	}
	// 	if s.HomeName != "" {
	// 		c.HomeName = s.HomeName
	// 	}
	// 	score := s.ColoredString()
	// 	for _, m := range msg.Args[0].Result.Data[0].EventMsgs {
	// 		if c.LastTime == 0 || c.LastTime < m.Content.T {
	// 			fmt.Fprintf(color.Output, "%s %s %s %s| %s\n", c.AwayName, score, c.HomeName, s.Process, m.String())
	// 			c.LastTime = m.Content.T
	// 		}
	// 	}
	// }
	fmt.Fprintf(color.Output, "%s\n", msg.Content)
}

func (c *Client) Start() {
	fmt.Println("----- 正在连接直播间 -----")
	// 初始化
	// c.Pid = 617 // 先默认设定一个数值，之后更新
	// c.ErrorCh = make(chan interface{}, 1)
	// c.InterruptCh = make(chan os.Signal, 1)
	// c.OprCh = make(chan interface{})
	// c.Th = time.NewTicker(HupuApp.LIVE_HEART_BEAT_PERIOD * time.Second)
	// c.Connect()

	// signal.Notify(c.InterruptCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// defer c.WsConn.Close()

	// go c.OnMessage()
	// c.finalize()
	// c.WaitClose()
	// api.QueryLiveTextList()
	c.init()

	for i := 1; i < 5; i++ {
		matchTextMsgs, err := api.QueryLiveTextList(c.Match.MatchId, c.liveActivityKey, c.LastCommentId)
		if err != nil {
			logger.Error.Fatal(err)
		}
		for _, msg := range matchTextMsgs {
			c.HandleLiveMsg(msg)
			c.LastCommentId = msg.CommentId
		}
	}

}

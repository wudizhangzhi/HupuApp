package live

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"github.com/wudizhangzhi/HupuApp"
	"github.com/wudizhangzhi/HupuApp/gohupu/api"
	"github.com/wudizhangzhi/HupuApp/gohupu/logger"
	"github.com/wudizhangzhi/HupuApp/gohupu/message"
)

type Client struct {
	WsConn      *websocket.Conn
	Domain      string       // (外部接口获取的参数)
	Pid         int          // (内部接口获取的参数)
	Game        message.Game // 比赛(外部接口获取的参数)
	HomeName    string       // 主队名字
	AwayName    string       // 客队名字
	Connected   bool         // ws中是否已连接
	LastTime    int          // (内部接口获取，用于比对直播数据时间)
	Th          *time.Ticker
	ErrorCh     chan interface{} // 错误channel
	OprCh       chan interface{} // 操作通道?
	InterruptCh chan os.Signal   // 中断信号
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
	token = strings.Split(string(respBody), ":50:60")[0]
	return token, nil
}

func (c *Client) Connect() {
	token, err := c.FetchToken()
	if err != nil {
		panic(err)
	}
	url := "ws://" + c.Domain + fmt.Sprintf("/socket.io/1/websocket/%s/?client=%s&t=%d&type=1&background=false", token, api.HupuHttpobj.IMEI, time.Now().Unix())
	logger.Info.Printf("创建连接: %s\n", url)
	wsConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		logger.Error.Fatalf("websocket连接失败: %v\n", err)
		panic(err)
	}
	c.WsConn = wsConn
}

func (c *Client) Send(msg string) {
	logger.Info.Printf("发送: %s\n", msg)
	c.WsConn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func loadResponse(respMsg []byte) (interface{}, error) {
	switch string(respMsg) {
	case "1::":
		return &message.MsgOne{}, nil
	case "2::":
		return &message.MsgTwo{}, nil
	case "1::/nba_v1":
		return &message.MsgNBAStart{}, nil
	default:
		msg := message.WsMsg{}
		if len(respMsg) < 11 {
			logger.Error.Printf("收到的消息长度不足: %s\n", respMsg)
			return nil, fmt.Errorf("收到的消息长度不足: %s", respMsg)
		}
		if err := json.Unmarshal(respMsg[11:], &msg); err != nil {
			return nil, err
		}
		return &msg, nil
	}
}

func (c *Client) heartbeat() {
	logger.Info.Println("心跳~")
	c.WsConn.WriteMessage(websocket.TextMessage, []byte("2:::"))
	// c.Send("2:::")
}

func (c *Client) OnMessage() {
	logger.Info.Printf("开始监听")
	for {
		msgType, p, err := c.WsConn.ReadMessage()
		if err != nil {
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					logger.Error.Printf("error: %v", err)
					c.OprCh <- "close"
				}
				break
			}
			logger.Error.Printf("接收数据报错, 退出: %s\n", err)
			if c.ErrorCh != nil {
				c.ErrorCh <- err
			}
			break
		}
		txtMsg := p
		switch msgType {
		case websocket.TextMessage:
			//
		case websocket.BinaryMessage:
			// txtMsg, err = o.GzipDecode(message)
		}
		// 处理response
		msgResp, err := loadResponse(txtMsg)
		if len(txtMsg) < 100 {
			logger.Info.Printf("收到的消息: %s\n", txtMsg)
		}
		if err != nil {
			c.ErrorCh <- err
			break
		}

		switch msg := msgResp.(type) {
		case *message.MsgOne:
			c.Send(message.MsgRespMsgOne)
		case *message.MsgTwo:
			if c.Connected {
				c.Send(message.MsgRespMsgTwo)
			} else {
				c.Connected = true
				c.Send(message.MsgRespMsgTwoConnected)
			}

		case *message.MsgNBAStart:
			c.Send(fmt.Sprintf(message.MsgRespMsgNBAStart, c.Game.Gid, c.Pid))
		case *message.WsMsg: // 如果是直播消息, 处理
			c.HandleLiveMsg(msg)
			if HupuApp.InterfaceToStr(msg.Args[0].RoomLiveType) == "-1" {
				// 比赛结束
				fmt.Println("----- 直播结束了 -----")
				c.OprCh <- "finish"
			}
		}

	}
}

// 格式化直播消息
func (c *Client) HandleLiveMsg(msg *message.WsMsg) {
	if len(msg.Args) > 0 && len(msg.Args[0].Result.Data) > 0 {
		s := msg.Args[0].Result.Score
		if s.AwayName != "" {
			c.AwayName = s.AwayName
		}
		if s.HomeName != "" {
			c.HomeName = s.HomeName
		}
		score := s.ColoredString()
		for _, m := range msg.Args[0].Result.Data[0].EventMsgs {
			if c.LastTime == 0 || c.LastTime < m.Content.T {
				fmt.Fprintf(color.Output, "%s %s %s %s| %s\n", c.AwayName, score, c.HomeName, s.Process, m.String())
				c.LastTime = m.Content.T
			}
		}
	}
}

func (c *Client) finalize() {
OutLoop:
	for {
		select {
		case <-c.InterruptCh:
			// 中断
			break OutLoop
		case err := <-c.ErrorCh:
			// 错误
			logger.Error.Printf("报错！： %s\n", err)
			break OutLoop
		case op := <-c.OprCh:
			// 输入
			switch op {
			case "finish":
				break OutLoop
			case "close":
				break OutLoop
			}
		case <-c.Th.C: // 心跳
			c.heartbeat()
		}
	}
	close(c.ErrorCh)
	close(c.InterruptCh)
	close(c.OprCh)
	c.WsConn.Close()
}

func (c *Client) WaitClose() {
	fmt.Println("----- 按任意键退出 -----")
	keyboard.GetSingleKey()
}

func (c *Client) Start() {
	fmt.Println("----- 正在连接直播间 -----")
	// 初始化
	c.Pid = 617 // 先默认设定一个数值，之后更新
	c.ErrorCh = make(chan interface{}, 1)
	c.InterruptCh = make(chan os.Signal, 1)
	c.OprCh = make(chan interface{})
	c.Th = time.NewTicker(HupuApp.LIVE_HEART_BEAT_PERIOD * time.Second)
	c.Connect()

	signal.Notify(c.InterruptCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer c.WsConn.Close()

	go c.OnMessage()
	c.finalize()
	c.WaitClose()
}

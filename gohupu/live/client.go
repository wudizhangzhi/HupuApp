package live

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type LiveClient struct {
	WsConn  *websocket.Conn
	Domain  string
	Th      *time.Ticker
	ErrorCh chan interface{}
}

// 获取token
func (c *LiveClient) fetchToken() (string, error) {
	var token string
	client := &http.Client{}
	t := time.Now().Second()
	url := "http://" + c.Domain + "/socket.io/1/"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Print(err)
		return "", err
	}
	q := req.URL.Query()
	q.Add("client", "")
	q.Add("t", fmt.Sprint(t))
	q.Add("type", "1")
	q.Add("background", "false")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	buf := make([]byte, 0)
	resp.Body.Read(buf)
	token = strings.Split(string(buf), ":50:60")[0]
	return token, nil
}

func (c *LiveClient) Connect() {
	url := "wss://" + c.Domain + "/v1/recognize?model=en-US_BroadbandModel&access_token="
	log.Println("创建连接")
	wsConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		panic(err)
	}
	c.WsConn = wsConn
}

func init() {

}

func (c *LiveClient) OnMessage() {
	for {
		msgType, message, err := c.WsConn.ReadMessage()
		if err != nil {
			log.Printf("接收数据报错, 退出: %s", err)
			if c.ErrorCh != nil {
				c.ErrorCh <- err
			}
			break
		}
		txtMsg := message
		switch msgType {
		case websocket.TextMessage:
			//
		case websocket.BinaryMessage:
			// txtMsg, err = o.GzipDecode(message)
		}
		// 处理response
	}
}

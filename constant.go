package HupuApp

const (
	DefaultLiveWebSocketDomain = "61.174.11.224:3081"
	DefaultTac                 = "tac.csv"
	API_VERSION                = "7.1.15"
	HUPU_SALT                  = ""
	// 接口
	// 直播接口
	API_INIT             = "https://games.mobileapi.hupu.com/1/" + API_VERSION + "/status/init"
	API_GET_GAMES        = "https://games.mobileapi.hupu.com/1/" + API_VERSION + "/%s/getGames"
	API_GET_PLAY_BY_PLAY = "https://games.mobileapi.hupu.com/1/" + API_VERSION + "/room/getPlaybyplay"
)

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

var (
	ANDROID_USER_AGENT = []string{
		"Mozilla/5.0 (Android 5.1.1; Tablet; rv:48.0) Gecko/48.0 Firefox/48.0",
		"Mozilla/5.0 (Linux; Android 4.4.2; Phoenix 2 Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2768.59 Mobile Safari/537.36",
		"Mozilla/5.0 (Android 4.4.4; Mobile; rv:46.0) Gecko/46.0 Firefox/46.0",
		"Mozilla/5.0 (Linux; Android 6.0; 7045Y Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2755.4 Mobile Safari/537.36",
	}
)

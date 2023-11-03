package HupuApp

type Config struct {
	// 接口设置
	API_VERSION            string `yml:"API_VERSION"`
	HUPU_SALT              string `yml:"HUPU_SALT"`
	LIVE_HEART_BEAT_PERIOD int    `yml:"LIVE_HEART_BEAT_PERIOD"`
	// 基础设置
	LOG_FILE int `yml:"LOG_FILE"`
}

package logger

import (
	"io"
	"log"
	"os"

	"github.com/wudizhangzhi/HupuApp"
)

var (
	Info  *log.Logger
	Debug *log.Logger
	Error *log.Logger
)

func init() {
	//日志输出文件
	file, err := os.OpenFile(HupuApp.LOG_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Faild to open error logger file:", err)
	}
	//自定义日志格式
	Info = log.New(io.MultiWriter(file), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(io.MultiWriter(file), "Debug: ", log.Ldate|log.Ltime|log.Lshortfile)
}

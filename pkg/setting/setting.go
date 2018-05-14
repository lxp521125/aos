package setting

import (
	"os"
	"strings"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/graylog"
	"github.com/apex/log/handlers/multi"
	"github.com/apex/log/handlers/text"
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort               int
	ReadTimeout            time.Duration
	WriteTimeout           time.Duration
	ResourceVideoSrc       string // 视频批量调用地址\
	ResourceLectureNoteSrc string // 讲义批量调用地址\
	UserTokenSrc           string // user token 调用地址
	PageSize               int
	JwtSecret              string
)

const (
	confName      = "goada.gaodun.com"
	TimeLayOut    = "2006-01-02"
	TimeLayOutHIS = "2006-01-02 15:04:05"
)

var TimeTags = map[string]int{
	"59f802c656142c6cea734ed5": 1,
	"59f802c656142c6cea734ed6": 1,
	"59f802c656142c6cea734ed7": 1,
	"59f802c656142c6cea734ed8": 1,
	"59f802c656142c6cea734edd": 1,
	"5a32229f56142cfe942d5cd1": 1,
	"5a32229f56142cfe942d5cd2": 1,
	"5a32229f56142cfe942d5cd4": 1,
}

//var Logger *log.Entry
// TODO log 不应该在setting当中
var Logger *log.Entry

func LoadConfig() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		// log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	loadBase()
	loadServer()
	loadApp()
	// 设置log
	Logger = GrayLog()
}

func loadBase() {

}

var grayLog = getGrayLog()

func getGrayLog() log.Handler {
	Cfg, _ = ini.Load("conf/app.ini")
	graylogInfo, _ := Cfg.GetSection("log")
	e, err := graylog.New(graylogInfo.Key("LOG_UDP").MustString("udp://g02.graylog.gaodunwangxiao.com:5504"))
	if err != nil {
		return &l
	}

	return e
}

func GrayLog(newFields ...map[string]interface{}) *log.Entry {
	graylogInfo, _ := Cfg.GetSection("log")
	isShowConsole := graylogInfo.Key("IS_SHOW_CONSOLE").MustBool(false)
	if isShowConsole {
		t := text.New(os.Stderr)
		log.SetHandler(multi.New(grayLog, t))
	} else {
		log.SetHandler(multi.New(grayLog))
	}

	fields := make(log.Fields)
	grayFields := graylogInfo.Key("LOG_FIELDS").MustString("item:ginlog")
	grayFieldsArray := strings.Split(grayFields, ",")
	if len(grayFieldsArray) > 0 {
		for i := 0; i < len(grayFieldsArray); i++ {
			temp := strings.Split(grayFieldsArray[i], ":")
			if len(temp) > 1 {
				fields[string(temp[0])] = temp[1]
			}
		}
	}

	if newFields != nil {
		for k, v := range newFields[0] {
			fields[k] = v
		}
	}
	level := graylogInfo.Key("LOG_LEVEL").MustInt(-1)
	log.SetLevel(log.Level(level))
	return log.WithFields(fields)
}

func Logrus(newFields ...map[string]interface{}) *logrus.Entry {
	graylogInfo, _ := Cfg.GetSection("log")
	level := graylogInfo.Key("LOG_LEVEL").MustInt(-1)
	logrus.SetLevel(logrus.Level(level))
	fields := make(logrus.Fields)
	grayFields := graylogInfo.Key("LOG_FIELDS").MustString("item:ginlog")
	grayFieldsArray := strings.Split(grayFields, ",")
	if len(grayFieldsArray) > 0 {
		for i := 0; i < len(grayFieldsArray); i++ {
			temp := strings.Split(grayFieldsArray[i], ":")
			if len(temp) > 1 {
				fields[string(temp[0])] = temp[1]
			}
		}
	}
	if newFields != nil {
		for k, v := range newFields[0] {
			fields[k] = v
		}
	}
	return logrus.WithFields(fields)
}
func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		// Logger.Fatalf("Fail to get section 'server': %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

	apiSection, err := Cfg.GetSection("api")
	ResourceVideoSrc = apiSection.Key("RESOURCE_VIDEO_SRC").String()
	ResourceLectureNoteSrc = apiSection.Key("RESOURCE_LECTURENOTE_SRC").String()
	UserTokenSrc = apiSection.Key("USER_TOKEN_API").String()
}

func loadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		// Logger.Fatalf("Fail to get section 'app': %v", err)
	}

	// JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

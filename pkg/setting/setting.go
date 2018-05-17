package setting

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/graylog"
	"github.com/apex/log/handlers/multi"
	"github.com/apex/log/handlers/text"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	RunMode string

	HTTPPort               int
	ReadTimeout            time.Duration
	WriteTimeout           time.Duration
	ResourceVideoSrc       string // 视频批量调用地址\
	ResourceLectureNoteSrc string // 讲义批量调用地址\
	UserTokenSrc           string // user token 调用地址
	PageSize               int
	JwtSecret              string
	CONSUL_URL             string
	CONSUL_LIST_NAME       string
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
var grayLog log.Handler

func LoadConfig() {

	loadBase()
	loadOTher()
	loadServer()
	loadApp()
	// 设置log
	grayLog = getGrayLog()
	Logger = GrayLog()
}

func loadBase() {
	cmdRoot := "app"
	viper.SetEnvPrefix(cmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName(cmdRoot)
	viper.AddConfigPath("./conf")
	viper.SetConfigType("yaml")
	// gopath := os.Getenv("GOPATH")
	// for _, p := range filepath.SplitList(gopath) {
	// 	peerpath := filepath.Join(p, "src/vip")
	// 	viper.AddConfigPath(peerpath)
	// }
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Println(fmt.Errorf("Fatal error when reading %s config file: %s\n", cmdRoot, err))
	}
}

func getGrayLog() log.Handler {
	e, err := graylog.New(viper.GetString("log.LOG_UDP"))
	if err != nil {
		return &l
	}
	return e
}

func GrayLog(newFields ...map[string]interface{}) *log.Entry {
	isShowConsole := viper.GetBool("log.IS_SHOW_CONSOLE")
	if isShowConsole {
		t := text.New(os.Stderr)
		log.SetHandler(multi.New(grayLog, t))
	} else {
		log.SetHandler(multi.New(grayLog))
	}

	fields := make(log.Fields)
	grayFields := viper.GetString("log.LOG_FIELDS")
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
	level := viper.GetInt("log.LOG_LEVEL")
	log.SetLevel(log.Level(level))
	return log.WithFields(fields)
}

func Logrus(newFields ...map[string]interface{}) *logrus.Entry {
	level := viper.GetInt("log.LOG_LEVEL")
	logrus.SetLevel(logrus.Level(level))
	fields := make(logrus.Fields)
	grayFields := viper.GetString("log.LOG_FIELDS")
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
	if viper.IsSet("RUN_MODE") {
		RunMode = viper.GetString("RUN_MODE")
	} else {
		RunMode = "debug"
	}

	if viper.IsSet("server.HTTP_PORT") {
		HTTPPort = viper.GetInt("server.HTTP_PORT")
	} else {
		HTTPPort = 6001
	}
	if viper.IsSet("server.READ_TIMEOUT") {
		ReadTimeout = viper.GetDuration("server.READ_TIMEOUT")
	} else {
		ReadTimeout = 60
	}

	if viper.IsSet("server.WRITE_TIMEOUT") {
		WriteTimeout = viper.GetDuration("server.WRITE_TIMEOUT")
	} else {
		WriteTimeout = 60
	}
}

func loadApp() {
	if viper.IsSet("app.PAGE_SIZE") {
		PageSize = viper.GetInt("app.PAGE_SIZE")
	} else {
		PageSize = 10
	}
}

func loadOTher() {
	if viper.IsSet("consul.CONSUL_URL") {
		CONSUL_URL = viper.GetString("consul.CONSUL_URL")
	} else {
		CONSUL_URL = ""
	}
	if viper.IsSet("consul.CONSUL_LIST_NAME") {
		CONSUL_LIST_NAME = viper.GetString("consul.CONSUL_LIST_NAME")
	} else {
		CONSUL_LIST_NAME = ""
	}
}

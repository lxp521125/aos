package graylog

// graylog
import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"sync"

	"fmt"
	"github.com/chenxuey/graylog-golang"
	"time"
)

var (
	GRAY_ITEM string
	p         sync.RWMutex
	GrayCount int = 10
)

// 获取主机名
func getHostName() string {
	name, err := os.Hostname()

	if err != nil {
		name = "localhost"
	}
	return name
}

func OpenGray() {
	for {
		t := time.NewTicker(15 * time.Minute)
		select {
		case <-t.C:
			if GrayCount <= 0 {
				GrayCount = 10
			}
		}
	}
}

// 定时15分钟解锁 gray log
func init() {
	go OpenGray()
}

// 写入日志
// param log string 日志内容
// param item map[string]interface{}
func GdLog(log string, item map[string]interface{}) {
	if GrayCount > 0 {
		host := getHostName()
		p.Lock()
		item["host"] = host
		if _, ok := item["item"]; !ok {
			item["item"] = "go-graylog-null"
		}
		item["short_message"] = log
		p.Unlock()

		g := gelf.New(gelf.Config{
			GraylogPort:     5504,
			GraylogHostname: "g02.graylog.gaodunwangxiao.com",
			//MaxChunkSizeWan: 42,
			//MaxChunkSizeLan: 1337,
		})

		logs, err := json.Marshal(&item)
		if err != nil {
			panic(err)
		}

		err = g.Log(string(logs))
		if err != nil {
			GrayCount--
			fmt.Println(err)
		}
	}
}

//AddGrayLog 添加日志
func AddGrayLog(logInfo ...interface{}) {
	var logStr string
	m := make(map[string]interface{})
	if len(logInfo) == 1 {
		logStr = logInfo[0].(string)
	} else {
		logStr = logInfo[0].(string)
		m = logInfo[1].(map[string]interface{})
	}
	p.Lock()
	m["item"] = GRAY_ITEM
	p.Unlock()
	GdLog(logStr, m)
}

//GrayXormSql 实现 xrom 打印日志接口
type GrayXormSql struct {
}

func (GrayXormSql) Write(p []byte) (n int, err error) {
	// todo sql 打印到 gray log
	if strings.Contains(string(p), "[sql]") {
		s := string(p)
		t := strings.Split(s, "took:")
		if len(t) > 1 {
			tst := strings.Trim(t[1], " ")
			tst = strings.Replace(tst, "ms\n", "", -1)
			sqlTime, _ := strconv.ParseFloat(tst, 10)
			m := map[string]interface{}{
				"user_sql":           s,
				"user_sql_exec_time": sqlTime,
			}
			AddGrayLog(s, m)
		} else {
			m := map[string]interface{}{
				"user_sql": s,
			}
			AddGrayLog(s, m)
		}

	}
	return 0, nil
}

//
//func main()  {
//	var w  sync.WaitGroup
//	for i:=1; i< 10; i++ {
//		GdLog("test542aa", map[string]interface{}{"item":"test-go", "leve": i})
//	}
//
//	w.Wait()
//}

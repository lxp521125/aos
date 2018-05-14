# godep 使用
```
安装 go get github.com/tools/godep

godep save 将项目中使用到的第三方库复制到项目的vendor目录下

godep restore godep会按照Godeps/Godeps.json内列表，依次执行go get -d -v 来下载对应依赖包到GOPATH路径下

说明：关于无法安装golang.org下的库时， You can also manually git clone the repository to $GOPATH/src/golang.org/x/sys
```

# 加上Swagger
```
安装 swag
$ go get -u github.com/swaggo/swag/cmd/swag
$ swag -v
依赖golang.org的包
如若无科学上网，可用以下

$ gopm get -g -v github.com/swaggo/swag/cmd/swag
$ cd $GOPATH/src/github.com/swaggo/swag/cmd/swag
$ go install

gopm 安装：$ go get -u github.com/gpmgo/gopm
```
```
安装 gin-swagger
$ go get -u github.com/swaggo/gin-swagger

$ go get -u github.com/swaggo/gin-swagger/swaggerFiles

使用
$ cd $GOPATH/src/aos
swag init 
地址：
http://127.0.0.1:6001/swagger/index.html
```

# Logger 使用
```
暂时支持 graylog
配置conf/app.ini 的log配置 LOG_FIELDS：打印到 graylog 的查询字段
Level = enum [-1,0,1,2,3,4] => ["all","debug","info","warn","error","fatal"]

引入 "aos/pkg/setting"
eg:（Debug、Info、Warn、Error、Fatal）、（Debugf、Infof、Warnf、Errorf、Fatalf）
setting.Logger.Info("string 类型")
setting.Logger.Info("string 类型",interface{}")

说明：setting.Logger 会得到一个grayLog的实例，后期会支持app.ini的参数配置，得到不同的实例,不需要额外的字段，可使用setting.Logger.WithField()生成实例
```

# Code码使用
```
"aos/pkg/errors"
eg:
errors.SYSERR // code码
errors.GetInfo()[errors.SYSERR] // code 对应的值
TODO:进度封装，方便使用
```

# GD Consul 使用
```
go get -u -x gitlab.gaodun.com/golib/consul
import 	"gitlab.gaodun.com/golib/consul"
consulData, _ := consul.GetConf("")
host := consulData["PUBLIC_MYSQL_DB_HOST"]
```
# Redis 使用
```Go
import (
	"aos/pkg/utils"
)
utils.RedisHandle.SetData("test1", "hhhhh", 0)
utils.RedisHandle.GetData("test1")
说明：现在只封装了SetData 和 GetData ，异常未处理，未打印到graylog中去
```

# Http Request Client 使用
```Go
import (
    "aos/pkg/utils"
)
data, err := utils.HttpHandle.Post("url", param, header)
data, err := utils.HttpHandle.Get("url", param, header)
详见：aos/pkg/utils/http.go

上传文件：
import (
	"github.com/imroc/req"
)
file, _ := os.Open("imroc.png")
req.Post(url, req.FileUpload{
	File:      file,
	FieldName: "file",       // FieldName 是表单字段名
	FileName:  "avatar.png", // Filename 是要上传的文件的名称，我们使用它来猜测mimetype，并将其上传到服务器上
})
使用req.UploadProgress监听上传进度

progress := func(current, total int64) {
	fmt.Println(float32(current)/float32(total)*100, "%")
}
req.Post(url, req.File("/Users/roc/Pictures/*.png"), req.UploadProgress(progress))
fmt.Println("upload complete")
```

# Mysql 使用
```Go
参见:"aos/persistence/demo.go"
类的结构
import (
	"aos/pkg/utils"
	"github.com/go-xorm/xorm"
)

// GdSubject 科目表
type GdSubject struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// GdSubjectFac 工厂
type GdSubjectFac struct {
	Table        GdSubject
	engine       *xorm.Engine
	RowsSlicePtr []GdSubject
}

// NewGdSubjectFacFac 初始化
func NewGdSubjectFacFac() *GdSubjectFac {
	var fac GdSubjectFac
	fac.engine, _ = utils.InitEng(0)
	return &fac
}

// FindAll 查询所有
func (myM *GdSubjectFac) FindAll(where string) error {
	err := myM.engine.Where(where).Find(&(myM.RowsSlicePtr))
	if err != nil {
        //
	}
	return err
}

使用
var subjectModel = persistence.NewGdSubjectFacFac()
    subjectModel.FindAll("")
subjectModel.RowsSlicePtr // 结果集

参考文献：https://www.kancloud.cn/kancloud/xorm-manual-zh-cn/56013
说明：选用xorm作为mysql的ORM引擎
```

# Sentry
```GO
package main

import (
	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
)

func init() {
	raven.SetDSN("https://<key>:<secret>@app.getsentry.com/<project>")
}

func main() {
	r := gin.Default()
	r.Use(sentry.Recovery(raven.DefaultClient, false))
	// only send crash reporting
	// r.Use(sentry.Recovery(raven.DefaultClient, true))
	r.Run(":8080")
}
```

# TODO list
- [x] Panic 处理
- [x] Sentry 日志
- [x] 加上Swagger
- [x] 支持Cors处理
- [x] SQL 驱动与ORM选取
- [x] SQL 日志打印到graylog
- [x] 输出数据打印到graylog
- [x] Http请求Clent
- [x] Session
- [x] X-Response-ID
- [x] Consul 读取
- [x] Redis 简单封装
- [x] 状态码统一管理
- [x] DDD设计实现,已经简单实现
- [ ] 表单验证
- [x] 支持各个项目生产


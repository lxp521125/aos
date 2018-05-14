package controller

import (
	"aos/pkg/errors"
	"aos/pkg/setting"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type ResponseObject struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// Base 基类
type Base struct {
}

func (bs *Base) ServerJSON(c *gin.Context, v interface{}, errorCode int) {
	if errorCode == 0 {
		bs.ServerJSONSuccess(c, v)
	} else {
		bs.ServerJSONError(c, v, errorCode)
	}
}

//ServerJSONSuccess 服务器返回
func (bs *Base) ServerJSONSuccess(c *gin.Context, v interface{}) {
	var rep ResponseObject
	rep.Code = errors.SUCCESSSTATUS
	rep.Message = errors.INFO[rep.Code]
	rep.Result = ""
	if v != nil {
		rep.Result = v
	}
	if b, jerr := json.Marshal(&rep); jerr == nil {
		setting.GinLogger(c).Infof("response result:", string(b))
	} else {
		setting.GinLogger(c).Infof("response result:", rep)
	}
	c.JSON(200, rep)
}

//ServerJSONError 服务器返回
func (bs *Base) ServerJSONError(c *gin.Context, v interface{}, errorCode int) {
	var rep ResponseObject
	rep.Code = errorCode
	rep.Message = errors.INFO[errorCode]
	rep.Result = ""
	if v != nil {
		rep.Result = v
	}
	if b, jerr := json.Marshal(&rep); jerr == nil {
		setting.GinLogger(c).Infof("response result:", string(b))
	} else {
		setting.GinLogger(c).Infof("response result:", rep)
	}
	c.JSON(200, rep)
}

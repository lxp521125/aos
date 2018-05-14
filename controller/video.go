package controller

import (
	"aos/pkg/errors"
	"aos/pkg/setting"
	"aos/pkg/utils"
	"aos/project/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoApi struct {
	Base
	Service *service.VideoFactory
}

func NewVideoController(service *service.VideoFactory) *VideoApi {
	return &VideoApi{Service: service}
}

func (this *VideoApi) Batch(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		this.ServerJSONError(c, err, errors.NODATA)
		return
	}

	// 获取uid
	uid, err := this.Service.GetUid(c)
	if err != nil {
		this.ServerJSONError(c, err, errors.NODATA)
		return
	}

	// 获取video信息
	videos, err := this.Service.Batch(fileHeader)
	if err != nil {
		this.ServerJSONError(c, err, errors.NODATA)
		return
	}

	for _, v := range videos {
		v.Creator_id = uid
		v.Partner_id = 1
	}

	if err != nil {
		this.ServerJSONError(c, err, errors.NOMOREDATA)
		return
	}

	result, err := utils.HttpHandle.PostBodyJson(setting.ResourceVideoSrc, videos)
	if err != nil {
		this.ServerJSONError(c, err, errors.NOMOREDATA)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (this *VideoApi) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

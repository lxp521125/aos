package controller

import (
	"aos/pkg/errors"
	"aos/pkg/setting"
	"aos/pkg/utils"
	"aos/project/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type LectureNoteApi struct {
	Base
	Service *service.LectureNoteFactory
}

func NewLectureNoteController(service *service.LectureNoteFactory) *LectureNoteApi {
	return &LectureNoteApi{Service: service}
}

func (this *LectureNoteApi) Batch(c *gin.Context) {
	fileHeader, err := c.FormFile("file")

	// 获取uid
	uid, err := this.Service.GetUid(c)
	if err != nil {
		this.ServerJSONError(c, err, errors.NODATA)
		return
	}

	if err != nil {
		this.ServerJSONError(c, err, errors.NOMOREDATA)
		return
	}
	lectureNotes, err := this.Service.Batch(fileHeader)
	if err != nil {
		this.ServerJSONError(c, err, errors.NOMOREDATA)
		return
	}
	for _, v := range lectureNotes {
		v.Creator_id = uid
		v.Partner_id = 1
		v.Path = "demo.pdf"
	}

	result, err := utils.HttpHandle.PostBodyJson(setting.ResourceLectureNoteSrc, lectureNotes)
	if err != nil {
		this.ServerJSONError(c, err, errors.NOMOREDATA)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (this *LectureNoteApi) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

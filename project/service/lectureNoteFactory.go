package service

import (
	"aos/pkg/setting"
	"aos/project/domain"
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type LectureNoteFactory struct {
	lectureNoteServer LectureNoticeService
	tagService        TagService
	userService       UserService
}

func NewLectureNoteFactory(lectureNoteServer LectureNoticeService, tagService TagService, userService UserService) *LectureNoteFactory {
	return &LectureNoteFactory{
		lectureNoteServer: lectureNoteServer,
		tagService:        tagService,
		userService:       userService,
	}
}

func (this *LectureNoteFactory) GetUid(c *gin.Context) (int64, error) {
	// 如果是测试环境则
	if setting.RunMode == "debug" {
		return 1, nil
	}

	token, ok := c.Get("token")
	if !ok {
		return -1, errors.New("无token 请求失败")
	}
	return this.userService.GetUserUid(fmt.Sprintf("%s", token))
}

func (this *LectureNoteFactory) Batch(file *multipart.FileHeader) ([]*domain.LectureNote, error) {
	lectureNotes, err := this.lectureNoteServer.Batch(file)
	if err != nil {
		return nil, err
	}
	// init 缓存
	cache := make(map[string]map[string]int64)

	for _, lectureNote := range lectureNotes {
		kind, key := lectureNote.GetTypeAndKey()
		if len(cache[kind]) == 0 {
			cache[kind] = make(map[string]int64)
		}

		// 查找到缓存的tag id则直接赋值
		id, ok := cache[kind][key]
		if ok {
			lectureNote.Tag_id = id
			continue
		}

		// 查找tag id
		tag, err := this.tagService.GetTag(kind, key)
		if err != nil {
			return nil, err
		}
		cache[kind][key] = tag.Id

		// 加入缓存当中
		lectureNote.Tag_id = id
	}
	return lectureNotes, nil
}

package service

import (
	"aos/pkg/setting"
	"aos/project/domain"
	"errors"
	"mime/multipart"

	"fmt"

	"github.com/gin-gonic/gin"
)

type VideoFactory struct {
	videoServer VideoServer
	tagService  TagService
	userService UserService
}

func NewVideoFactory(videoService VideoServer, tagService TagService, userService UserService) *VideoFactory {
	return &VideoFactory{
		videoServer: videoService,
		tagService:  tagService,
		userService: userService,
	}
}

func (this *VideoFactory) GetUid(c *gin.Context) (int64, error) {
	// 如果是测试环境则
	if setting.RunMode == "debug" {
		return 1, nil
	}

	token, ok := c.Get("token")
	if !ok {
		return -1, errors.New("无token 请求失败")
	}

	return this.userService.GetUserUid(fmt.Sprint("%s", token))
}

func (this *VideoFactory) Batch(file *multipart.FileHeader) ([]*domain.Video, error) {
	videos, err := this.videoServer.Batch(file)
	if err != nil {
		return nil, err
	}

	// init 缓存
	cache := make(map[string]map[string]int64)

	for _, video := range videos {
		kind, key := video.GetTypeAndKey()
		if len(cache[kind]) == 0 {
			cache[kind] = make(map[string]int64)
		}

		// 查找到缓存的tag id则直接赋值
		id, ok := cache[kind][key]
		if ok {
			video.Tag_id = id
			continue
		}

		// 查找tag id
		tag, err := this.tagService.GetTag(kind, key)
		if err != nil {
			return nil, err
		}
		cache[kind][key] = tag.Id

		// 加入缓存当中
		video.Tag_id = tag.Id
	}
	return videos, nil
}

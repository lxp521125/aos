package service

import (
	"aos/project/domain"
	"errors"
	"strings"

	"mime/multipart"

	"aos/pkg/utils"

	"github.com/360EntSecGroup-Skylar/excelize"
)

const (
	SUFFIX = "xlsx"
)

type VideoServiceImpl struct {
}

func NewVideoServiceImpl() *VideoServiceImpl {
	return &VideoServiceImpl{}
}
func (this *VideoServiceImpl) Batch(fileHeader *multipart.FileHeader) ([]*domain.Video, error) {
	// filter
	fileName := fileHeader.Filename
	tmp := strings.Split(fileName, ".")
	if len(tmp) < 2 {
		return nil, errors.New("传入file不存在")
	}

	// validate file is .xlsx
	if tmp[len(tmp)-1] != SUFFIX {
		return nil, errors.New("传入的file格式不正确")
	}

	file, err := fileHeader.Open()
	defer file.Close()
	if err != nil {
		return nil, err
	}

	// read excel
	excel, err := excelize.OpenReader(file)
	if err != nil {
		return nil, err
	}
	sheet := excel.GetRows("Sheet1")
	video := &domain.Video{}

	result, err := utils.ReadExcel(sheet, video, 3, 0)
	if err != nil {
		return nil, err
	}
	var videos []*domain.Video
	for _, v := range result {
		videos = append(videos, v.(*domain.Video))
	}
	return videos, nil
}

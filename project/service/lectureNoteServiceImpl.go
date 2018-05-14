package service

import (
	"aos/pkg/utils"
	"aos/project/domain"
	"errors"
	"mime/multipart"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type LectureNoteServiceImpl struct {
}

func NewLectureNoteServiceImpl() *LectureNoteServiceImpl {
	return &LectureNoteServiceImpl{}
}
func (this *LectureNoteServiceImpl) Batch(fileHeader *multipart.FileHeader) ([]*domain.LectureNote, error) {
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
	video := &domain.LectureNote{}

	result, err := utils.ReadExcel(sheet, video, 3, 0)
	if err != nil {
		return nil, err
	}
	var lectureNote []*domain.LectureNote
	for _, v := range result {
		lectureNote = append(lectureNote, v.(*domain.LectureNote))
	}
	return lectureNote, nil
}

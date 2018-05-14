package service

import (
	"aos/project/domain"
	"mime/multipart"
)

type LectureNoticeService interface {
	Batch(file *multipart.FileHeader) ([]*domain.LectureNote, error) // 批量
}

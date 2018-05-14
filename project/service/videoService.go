package service

import (
	"aos/project/domain"
	"mime/multipart"
)

type VideoServer interface {
	Batch(file *multipart.FileHeader) ([]*domain.Video, error) // 批量
}

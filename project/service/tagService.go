package service

import "aos/project/domain"

type TagService interface {
	GetTag(tagType, tagName string) (domain.TagModel, error)
}

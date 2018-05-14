package service

import "aos/project/domain"

type TagServiceImpl struct {
	TagDao domain.TagDAO //要使用的接口起名字
}

func NewTagServiceImpl(tagDao domain.TagDAO) *TagServiceImpl {
	return &TagServiceImpl{
		TagDao: tagDao,
	}
}

func (this *TagServiceImpl) GetTag(tagType, tagName string) (domain.TagModel, error) {
	return this.TagDao.GetTag(tagType, tagName)
}

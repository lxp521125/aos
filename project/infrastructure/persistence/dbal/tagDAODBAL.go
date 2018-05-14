package persistence

import (
	"aos/project/domain"

	"github.com/go-xorm/xorm"
)

type TagDAODBAL struct {
	Table  domain.TagModel
	Engine *xorm.Engine
}

func NewTagDAOBAL(engine *xorm.Engine) *TagDAODBAL {
	return &TagDAODBAL{
		Engine: engine,
	}
}

func (this TagDAODBAL) GetTag(tagType, tagName string) (domain.TagModel, error) {
	tag := this.Table
	ok, err := this.Engine.Table("resource_tags").Where("name = ? and payload_type = ?", tagName, tagType).Get(&tag)
	if err != nil || !ok {
		return tag, err
	}
	return tag, nil
}

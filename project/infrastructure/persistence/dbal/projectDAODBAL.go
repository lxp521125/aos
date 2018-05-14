package persistence

import (
	"aos/project/domain"

	"github.com/go-xorm/xorm"
)

type ProjectDAODBAL struct {
	Table  domain.ProjectModel
	Engine *xorm.Engine
}

func (c ProjectDAODBAL) List(b int) domain.ProjectModel {
	bp := c.Table
	c.Engine.Table("gd_project").Where("id = 8").Get(&bp)
	return bp
}

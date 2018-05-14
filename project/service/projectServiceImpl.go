package service

import (
	"aos/project/domain"
	"fmt"
)

type ProjectServiceImpl struct { //接口的定义
	// ProjectService                   //实现的接口不用起名字 ?重复的调用了？不写也是能实现继承
	ProjectDAO domain.ProjectDAO //要使用的接口起名字
}

func NewProjectServiceImpl(projectDAO domain.ProjectDAO) *ProjectServiceImpl {
	return &ProjectServiceImpl{projectDAO}
}

func (c *ProjectServiceImpl) List(projectId int) domain.ProjectModel {
	fmt.Println("come")
	return c.ProjectDAO.List(projectId)
}

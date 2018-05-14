package service

import (
	"aos/project/domain"
)

type ProjectService interface {
	List(projectId int) domain.ProjectModel
}

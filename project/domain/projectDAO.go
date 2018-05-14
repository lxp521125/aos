package domain

type ProjectDAO interface {
	List(projectId int) ProjectModel
}

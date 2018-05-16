package container

import (
	"aos/controller"
	"aos/pkg/utils"
	"aos/project/domain"
	"aos/project/infrastructure/persistence/dbal"
	"aos/project/service"

	"github.com/go-xorm/xorm"
)

var containerInstance *Container

func GetContainer() *Container {
	if containerInstance == nil {
		containerInstance = &Container{}
		containerInstance.init()
	}
	return containerInstance
}

type Container struct {
}

func (this *Container) init() {
	//var dbConnect *xorm.Engine = initEng(0)
}

func initEng(num int) *xorm.Engine {
	eng, _ := utils.InitEng(num)
	return eng
}
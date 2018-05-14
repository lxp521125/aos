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
	TestApi        *controller.TestApi
	VideoApi       *controller.VideoApi
	LectureNoteApi *controller.LectureNoteApi
}

func (this *Container) init() {
	var dbConnect *xorm.Engine = initEng(0)
	var projectDAO domain.ProjectDAO = NewProjectDAODBAL(dbConnect)

	var ProjectService service.ProjectService = service.NewProjectServiceImpl(projectDAO)
	this.TestApi = controller.NewProjectController(ProjectService)

	//user server init
	var userService service.UserService = service.NewUserServiceImpl()
	// video server init
	var videoService service.VideoServer = service.NewVideoServiceImpl()
	var tagdao domain.TagDAO = persistence.NewTagDAOBAL(dbConnect)

	var tagservice service.TagService = service.NewTagServiceImpl(tagdao)
	var videoFactory = service.NewVideoFactory(videoService, tagservice, userService)
	this.VideoApi = controller.NewVideoController(videoFactory)

	// lecture notice server init
	var lectureNoteService service.LectureNoticeService = service.NewLectureNoteServiceImpl()
	var lectureNoteFactory = service.NewLectureNoteFactory(lectureNoteService, tagservice, userService)
	this.LectureNoteApi = controller.NewLectureNoteController(lectureNoteFactory)

}

func initEng(num int) *xorm.Engine {
	eng, _ := utils.InitEng(num)
	return eng
}

//方法可以写在这里也可以放在具体的内部 采用哪个？
func NewProjectDAODBAL(eng *xorm.Engine) *persistence.ProjectDAODBAL {
	var a *persistence.ProjectDAODBAL = &persistence.ProjectDAODBAL{}
	a.Engine = eng
	return a
}

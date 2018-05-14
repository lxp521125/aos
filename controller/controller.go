package controller

import (
	"aos/pkg/errors"
	"aos/pkg/utils"
	"aos/project/service"
	"fmt"

	"aos/persistence"
	"aos/secret"

	"github.com/gin-gonic/gin"
	"aos/pkg/setting"
)

type TestApi struct {
	Base
	ProjectService service.ProjectService
}

func NewProjectController(projectService service.ProjectService) *TestApi {
	var c = &TestApi{}
	c.ProjectService = projectService
	return c
}

func AddNewSecret(c *gin.Context) {

}

// @Summary 获取S
// @Produce  json
// @Param access_key path string true "秘钥KEY"
// @Success 200 {string} json "{"status": 1,"message": "","result": {"access_key": "xxx","access_secret": ""}}"
// @Router /secret/{access_key} [get]
func (myc *TestApi) GetS(c *gin.Context) {

	// utils.RedisHandle.SetData("test1", "hhhhh", 0)
	// utils.HttpHandle.Debug = true
	data, err := utils.HttpHandle.Get("http://t-goada.gaodun.com/homework/oneStatisticsInfo?homework_id=33&is_complete_all=1", nil, nil)
	fmt.Println(utils.RedisHandle.GetData("test1"))
	fmt.Println(err)
	// authentication := CreateSecretFromRequest(c)
	// fmt.Println("Access key is " + authentication.AccessKey)
	// fmt.Println("Access secret is " + authentication.AccessSecret)
	// _, err := secretServiceFacade.Authenticate(authentication)
	// if nil != err {
	// 	fmt.Println(err)
	// }
	myc.ServerJSON(c, data, errors.SUCCESSSTATUS)
	// c.JSON(200, ResponseObject{
	// 	1,
	// 	"",
	// 	authentication,
	// })
}

func (myc *TestApi) GetDbTest(c *gin.Context) {
	var subjectModel = persistence.NewGdSubjectFacFac()
	subjectModel.FindAll("")
	myc.ServerJSON(c, subjectModel.RowsSlicePtr, errors.SUCCESSSTATUS)
}

func CreateSecretFromRequest(c *gin.Context) secret.Secret {
	accessKey := c.PostForm("access_key")
	if accessKey == "" {
		accessKey = c.Param("access_key")
	}
	accessSecret := c.DefaultQuery("access_secret", "")

	return secret.Secret{
		accessKey,
		accessSecret,
	}
}

func (myc *TestApi) GetServiceTest(c *gin.Context) {
	// println(reflect.TypeOf(new(project_service_impl.ProjectServiceStruct)))

	data := myc.ProjectService.List(8)

	myc.ServerJSON(c, data, errors.SUCCESSSTATUS)

}

func (myc *TestApi) TestGraylog(c *gin.Context) {
	// println(reflect.TypeOf(new(project_service_impl.ProjectServiceStruct)))
	setting.GinLogger(c).Info("abc")

}

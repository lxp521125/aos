package test

import (
	"aos/pkg/setting"
	"aos/pkg/utils"
	"aos/routers"
	"os"
	"testing"

	"net/http/httptest"

	"net/http"

	"net/url"

	"strings"

	"aos/project/infrastructure/persistence/dbal"

	"fmt"

	"aos/project/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var r *gin.Engine
var w *httptest.ResponseRecorder

func init() {
	setting.LoadConfig()
	// init db
	if err := utils.InitEngine(); err != nil {
		panic(err)
		os.Exit(0)
	}
	r = routers.InitRouter()
	w = httptest.NewRecorder()
}

func TestPingHandleVideoController(t *testing.T) {
	var postData = url.Values{}
	data := postData.Encode()
	req, _ := http.NewRequest("POST", "/v1/ping", strings.NewReader(data))
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestBatchHandleVideoController(t *testing.T) {
	//var postData = url.Values{}
	//
	//postData["file"] =
}

func TestTag(t *testing.T) {
	dbConnect, err := utils.InitEng(0)
	if err != nil {
		t.Error(err)
	}
	tagDAODBAL := persistence.TagDAODBAL{
		Engine: dbConnect,
	}
	tag, err := tagDAODBAL.GetTag("subject", "F4")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(tag)
}

func TestToken(t *testing.T) {
	userServiceImpl := &service.UserServiceImpl{}
	token := "Basic eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoiVEVjVHY3VTdzcWNIR2p3clVxalBXc2JWVDNMZytsb05hYjZGQnpVNENHMFF6L2IwREFoVnJvWnBpMnp1bjZkRlFDSDE1dXJrZmlSaXppYWwyVEJQRE9zK1RnZ0svRDhGa1lPejQ2Q1RoMXdRRCtrVHAyalFGSFBDdFRzb0pPUTFtcW1YbHNZcGIxWnVvUlpsOUN2Y2RKaU0wMHZ5OVRXNzZBOHpsd2ZWWVBFSXFsWTlDRHUxTDJTNnJPUG9pMWhGTCtQZm1LSFF2b3E0MEFKOC93TkdnNVlJS0NTdlNZdDZrNGhNSWljbXZicU5GNWFyL2hsMlpTTkxsSW14MXNnSGhOY21RL2ZHOFozeDdNWnhENUdkRm02YXlhZ3NJb3ViaW1Bc2RaNkU2M1o1V2lsbks0R0pMbGJTVUEyN0k0UkkzZ0pUbC9QZy9iNmVidWZoK3djNXVRPT0iLCJleHAiOjE1MjU3NzU0NzAsImp0aSI6IjRmNGNkZTNiLTM3YTMtNGM4Yy1hYjAyLWZkOWY1YjRiMjhiYyIsImlhdCI6MTUyNTc2ODI3MCwiaXNzIjoiMTgwMTA3In0.scWkZv3lyNhZSultL1G8br6lin-UUzhVhmqNCiOBrUY"
	uid, err := userServiceImpl.GetUserUid(token)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(uid)
}

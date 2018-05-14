package service

import (
	"aos/pkg/setting"
	"aos/pkg/utils"
	"errors"
)

type UserServiceImpl struct {
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{}
}
func (this *UserServiceImpl) GetUserUid(token string) (int64, error) {
	header := make(map[string]string)
	header["origin"] = "gaodun.com"
	header["Authentication"] = token
	data, err := utils.HttpHandle.Post(setting.UserTokenSrc, nil, header)
	if err != nil {
		return -1, err
	}
	result := data.(map[string]interface{})
	resultMap := result["result"].(map[string]interface{})
	resultData := resultMap["UserData"]
	//resultData := resultMap[""]
	uid, ok := resultData.(map[string]interface{})["Uid"]
	if !ok {
		return -1, errors.New("token 返回值 无Uid key")
	}
	return int64(uid.(float64)), nil
}

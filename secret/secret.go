package secret

// 领域模型
type Secret struct {
	AccessKey    string `json:"access_key"`
	AccessSecret string `json:"access_secret"`
}

// 比较两个Secret是否相等
func (s *Secret) equal(authenticated Secret) bool {
	return authenticated.AccessSecret == s.AccessSecret &&
		authenticated.AccessKey == s.AccessKey
}

// Secret 工厂的定义
type SecretFactory interface {
	Create() Secret
}

// 安全服务的Facade接口
type SecretServiceFacade interface {
	Add(authenticated Secret) (*Secret, error)
	Authenticate(authenticated Secret) (bool, error)
	// TODO: 完成实现
	// Remove(secrete * Secret)error
}

// Secret的DAO接口
type SecretDAO interface {
	Add(secret Secret) error
	FindOne(secretAccessKey string) (*Secret, error)
	// TODO: 完成实现
	// Remove(secret * Secret)error
}

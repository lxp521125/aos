package secret

import (
	"math/rand"
	"time"
)

// Factory 的实现
// Factory 用于创建复杂对象或者和领域无关的模型
type SecretFactoryImpl struct {
}

func (f *SecretFactoryImpl) Create() Secret {
	return Secret{
		RandStringRunes(32),
		RandStringRunes(32),
	}
}

func NewSecretFactory() SecretFactory {
	return &SecretFactoryImpl{}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

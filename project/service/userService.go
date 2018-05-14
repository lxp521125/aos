package service

type UserService interface {
	GetUserUid(token string) (int64, error)
}

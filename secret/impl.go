package secret

import (
	"aos/pkg/errors"
	"fmt"
)

// Service facade
type SecretServiceFacadeImpl struct {
	secretDAO SecretDAO
	factory   SecretFactory
}

func (s *SecretServiceFacadeImpl) Add(authenticated Secret) (*Secret, error) {
	found, err := s.secretDAO.FindOne(authenticated.AccessKey)
	if nil != err {
		return nil, err
	}
	if found.equal(authenticated) {
		return nil, errors.New(1, "Not Authenticate")
	}
	newSecret := s.factory.Create()
	err = s.secretDAO.Add(newSecret)
	return &newSecret, err
}

func (s *SecretServiceFacadeImpl) Authenticate(authenticated Secret) (bool, error) {
	fmt.Println("SecretServiceFacadeImpl.Authenticate")
	found, err := s.secretDAO.FindOne(authenticated.AccessKey)

	if nil != err {
		return false, err
	}
	if found.AccessSecret != authenticated.AccessSecret {
		return false, nil
	}
	return true, err
}

func NewSecretServiceFacadeImpl(secretDAO SecretDAO, factory SecretFactory) SecretServiceFacade {
	return &SecretServiceFacadeImpl{
		secretDAO,
		factory,
	}
}

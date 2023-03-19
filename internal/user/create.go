package user

import (
	"errors"
)

var AlreadyExists = errors.New("createUserService: user already exists")

type CreateUserService struct {
	ur Repository
}

func (c *CreateUserService) Execute(ID string) (*User, error) {
	u, err := c.ur.FindById(ID)
	if err == nil {
		return nil, AlreadyExists
	}

	u = NewUser(ID)
	_ = c.ur.Save(u)

	return u, nil
}

func NewCreateUserService(ur Repository) *CreateUserService {
	return &CreateUserService{ur}
}

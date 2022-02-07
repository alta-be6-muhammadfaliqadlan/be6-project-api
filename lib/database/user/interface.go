package user

import (
	"part3/models/user"
)

type User interface {
	Create(newUser user.User) (user.User, error)
	// GetById(id int) (user.User, error)
	// UpdateById(id int, userReg request.UserRegister) (user.User, error)
	// DeleteById(id int) (gorm.DeletedAt, error)
	// GetAll() ([]response.UserResponse, error)
}

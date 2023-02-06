package repository

import (
	"github.com/deeprave/go-auth/models"
)

type DB interface {
	Close()

	GetUsers(activeOnly bool, window ...*Window) ([]*models.User, error)

	GetUserById(id int64) (*models.User, error)
	GetUserByName(name string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)

	GetCredentialsForUser(id int64) ([]*models.Credential, error)
}

package interfaces

import (
	"github.com/rikiya/go-clean/src/entity"
)

// UserRepository ...
type UserRepository interface {
	Store(entity.User) error
	Index() ([]entity.User, error)
	Update(int, entity.User) error
	Delete(id int) error
}

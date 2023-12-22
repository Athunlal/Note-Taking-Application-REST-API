package interfaces

import "github.com/athunlal/Note-Taking-Application/pkg/domain"

type UserRepo interface {
	CreateUser(userData *domain.User) error
	FindUserById(userId int) (*domain.User, error)
	FindUserByEmail(userData domain.User) (*domain.User, error)

	CreateNote(Notes domain.Notes) error
	GetNoteById(Id string) (*domain.Notes, error)
	GetNotes(userId int) ([]domain.Notes, error)
	DeleteNoteById(Id string) error
}

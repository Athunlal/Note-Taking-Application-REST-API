package interfaces

import "github.com/athunlal/Note-Taking-Application/pkg/domain"

type UserRepo interface {
	CreateUser(userData *domain.User) error
	FindUserByEmail(userData domain.User) (*domain.User, error)

	CreateNote(Notes domain.Notes) error
	GetNotes(sId string) ([]domain.Notes, error)
	DeleteNoteById(Id string) error
}

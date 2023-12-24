package interfaces

import "github.com/athunlal/Note-Taking-Application/pkg/domain"

type UserUseCase interface {
	RegisterUser(user domain.User) error
	UserLogin(user domain.User) (*domain.User, error)

	CreateNotes(note domain.Notes) error
	GetNotes(sid string) ([]domain.Notes, error)
	DeleteNote(notes domain.Notes) error
}

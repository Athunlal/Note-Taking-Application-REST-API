package interfaces

import "github.com/athunlal/Note-Taking-Application/pkg/domain"

type UserUseCase interface {
	RegisterUser(user domain.User) error
	UserLogin(user domain.User) (*domain.User, error)

	CreateNotes(note domain.Notes) error
	GetNote(notes domain.Notes) (*domain.Notes, error)
	GetNotes(id int) ([]domain.Notes, error)
	DeleteNote(notes domain.Notes) error
}

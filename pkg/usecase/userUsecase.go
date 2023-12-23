package usecase

import (
	"errors"

	"github.com/athunlal/Note-Taking-Application/pkg/domain"
	interfaces "github.com/athunlal/Note-Taking-Application/pkg/repository/interface"
	userCaseInterface "github.com/athunlal/Note-Taking-Application/pkg/usecase/interfaces"
)

type userUseCase struct {
	repo interfaces.UserRepo
}

// CreateNotes implements interfaces.UserUseCase.
func (use *userUseCase) CreateNotes(note domain.Notes) error {
	return use.repo.CreateNote(note)
}

// DeleteNote implements interfaces.UserUseCase.
func (use *userUseCase) DeleteNote(notes domain.Notes) error {
	return use.repo.DeleteNoteById(notes.Sid)
}

// GetNote implements interfaces.UserUseCase.
func (use *userUseCase) GetNote(notes domain.Notes) (*domain.Notes, error) {
	return use.repo.GetNoteById(notes.Sid)
}

// GetNotes implements interfaces.UserUseCase.
func (use *userUseCase) GetNotes(userId int) ([]domain.Notes, error) {
	return use.repo.GetNotes(userId)
}

// RegisterUser implements interfaces.UserUseCase.
func (use *userUseCase) RegisterUser(user domain.User) error {
	return use.repo.CreateUser(&user)
}

// UserLogin implements interfaces.UserUseCase.
func (use *userUseCase) UserLogin(user domain.User) (*domain.User, error) {
	res, err := use.repo.FindUserByEmail(user)
	if err != nil {
		return nil, err
	}
	if res.Username == "" {
		return nil, errors.New("User not found")
	}
	if res.Password != user.Password {
		return nil, errors.New("Login credintials fail")
	}
	return nil, nil
}

func NewUserUseCase(repo interfaces.UserRepo) userCaseInterface.UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

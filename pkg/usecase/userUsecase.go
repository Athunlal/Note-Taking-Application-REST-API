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
	if note.Note == "" || note.Sid == "" {
		return errors.New("Empty note")

	}
	return use.repo.CreateNote(note)
}

// DeleteNote implements interfaces.UserUseCase.
func (use *userUseCase) DeleteNote(notes domain.Notes) error {
	res, err := use.repo.GetNotes(notes.Sid)
	if err != nil {
		return err
	}
	if len(res) < 1 {
		return errors.New("not found")
	}

	err = use.repo.DeleteNoteById(notes.Sid)
	if err != nil {
		return err
	}
	return nil
}

// GetNotes implements interfaces.UserUseCase.
func (use *userUseCase) GetNotes(sId string) ([]domain.Notes, error) {
	res, err := use.repo.GetNotes(sId)
	if err != nil {
		return nil, err
	}
	if len(res) < 1 {
		return nil, errors.New("Not Found")
	}
	return res, nil
}

// RegisterUser implements interfaces.UserUseCase.
func (use *userUseCase) RegisterUser(user domain.User) error {
	res, err := use.repo.FindUserByEmail(user)
	if err != nil {
		if err := use.repo.CreateUser(&user); err != nil {
			return err
		}
		return nil
	}
	if res != nil && res.Email != "" {
		return errors.New("user with the given email already exists")
	}
	return nil
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

	return &domain.User{
		Model:    res.Model,
		Username: res.Username,
		Password: "",
		Email:    res.Email,
	}, nil
}

func NewUserUseCase(repo interfaces.UserRepo) userCaseInterface.UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

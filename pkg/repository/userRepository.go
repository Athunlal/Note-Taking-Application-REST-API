package repository

import (
	"github.com/athunlal/Note-Taking-Application/pkg/domain"
	interfaces "github.com/athunlal/Note-Taking-Application/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

// GetNotes implements interfaces.UserRepo.
func (db *userDatabase) GetNotes(sId string) ([]domain.Notes, error) {
	var notes []domain.Notes
	res := db.DB.Where("sid = ?", sId).Find(&notes)
	if res.Error != nil {
		return nil, res.Error
	}
	return notes, nil
}

// CreateNote implements interfaces.UserRepo.
func (db *userDatabase) CreateNote(Notes domain.Notes) error {
	res := db.DB.Create(&Notes)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// CreateUser implements interfaces.UserRepo.
func (db *userDatabase) CreateUser(userData *domain.User) error {
	res := db.DB.Create(&userData)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// DeleteNoteById implements interfaces.UserRepo.
func (db *userDatabase) DeleteNoteById(Id string) error {
	var note domain.Notes
	if err := db.DB.Where("sid = ?", Id).Delete(&note).Error; err != nil {
		return err
	}
	return nil
}

// FindUserByEmail implements interfaces.UserRepo.
func (db *userDatabase) FindUserByEmail(userData domain.User) (*domain.User, error) {
	res := db.DB.Where("email = ?", userData.Email).First(&userData)
	if res.Error != nil {
		return nil, res.Error
	}
	return &userData, nil
}

// FindUserById implements interfaces.UserRepo.
func (*userDatabase) FindUserById(userId int) (*domain.User, error) {
	panic("unimplemented")
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &userDatabase{
		DB: db,
	}
}

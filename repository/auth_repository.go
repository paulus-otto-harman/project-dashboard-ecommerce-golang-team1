package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"project/database"
	"project/domain"
)

type AuthRepository struct {
	db     *gorm.DB
	cacher database.Cacher
}

func NewAuthRepository(db *gorm.DB, cacher database.Cacher) *AuthRepository {
	return &AuthRepository{db: db, cacher: cacher}
}

func (repo AuthRepository) Authenticate(user domain.User) (string, string, bool, error) {
	var userFound bool
	if err := repo.db.Model(&domain.User{}).Select("count(*)>0").Where(user).Find(&userFound).Error; err != nil {
		return "", "", false, errors.New("invalid username and/or password")
	}

	if userFound {
		token := uuid.New().String()

		if err := repo.cacher.Set(user.Username, token); err != nil {
			return "", "", false, err
		}

		return user.Username, token, true, nil
	}

	return "", "", false, nil
}

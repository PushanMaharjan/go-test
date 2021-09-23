package repository

import (
	"go-fx-test/infrastructure"

	"gorm.io/gorm"
)

type UserRepository struct {
	infrastructure.Database
}

func NewUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{db}
}

func (r UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	r.Database.DB = trxHandle
	return r
}

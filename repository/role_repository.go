package repository

import (
	"go-fx-test/infrastructure"

	"gorm.io/gorm"
)

type RoleRepository struct {
	infrastructure.Database
}

func NewRoleRepository(db infrastructure.Database) RoleRepository {
	return RoleRepository{db}
}

func (r RoleRepository) WithTrx(trxHandle *gorm.DB) RoleRepository {
	r.Database.DB = trxHandle
	return r
}

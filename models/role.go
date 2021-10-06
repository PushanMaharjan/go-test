package models

import (
	"go-fx-test/lib"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	Base
	Role string `json:"role"`
}

type UpdateRoleInput struct {
	Role string `json:"role"`
}

func (u Role) TableName() string {
	return "roles"
}

func (t *Role) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	t.ID = lib.BinaryUUID(id)
	return err
}

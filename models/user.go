package models

import (
	"go-fx-test/lib"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Base
	Fname string `json:"fname" form:"fname"`
	Lname string `json:"lname" form:"lname"`
	Admin bool   `json:"admin" form:"admin"`
}

func (u User) TableName() string {
	return "users"
}

func (t *User) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	t.ID = lib.BinaryUUID(id)
	return err
}

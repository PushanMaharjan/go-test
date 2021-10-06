package models

import (
	"go-fx-test/lib"
)

type Base struct {
	ID lib.BinaryUUID `json:"id"`
}

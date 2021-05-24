package model

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
}

type ModelInterface interface {}
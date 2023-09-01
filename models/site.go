package models

import (
	"gorm.io/gorm"
)

type Site struct {
	gorm.Model
	Title       string
	Description string
	MainMenu    []string
}

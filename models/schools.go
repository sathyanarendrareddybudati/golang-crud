package models

import (
	"gorm.io/gorm"
)

type Schools struct {
	gorm.Model
	Name  string `gorm:"column:name;type:varchar(255);not null"`
	Place string `gorm:"column:place;type:varchar(255);not null"`
}

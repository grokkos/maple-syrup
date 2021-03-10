package models

import (
	_ "github.com/jinzhu/gorm"
)

type User struct {
	ID  uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Log int    `gorm:"not null" json:"log"`
}

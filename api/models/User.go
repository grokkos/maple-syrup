package models

import (
	_ "github.com/jinzhu/gorm"
)

type User struct {
	ID   uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null;unique" json:"name"`
}

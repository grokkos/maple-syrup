package models

import (
	_ "github.com/jinzhu/gorm"
)

type Batch struct {
	ID          uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Dispatched  bool   `gorm:"default:false"`
	Summary     int    `gorm:"not null" json:"summary"`
	BatchUserID uint32 `sql:"type:int REFERENCES users(id)" json:"batch_user_id"`
	BatchUser   User   `json:"batch_user"`
}

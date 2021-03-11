package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
)

type Batch struct {
	ID          uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Dispatched  bool   `gorm:"default:false"`
	Summary     int    `gorm:"not null" json:"summary"`
	BatchUserID uint32 `sql:"type:int REFERENCES users(id)" json:"batch_user_id"`
	BatchUser   User   `json:"batch_user"`
}

func (u *Batch) FindAllBatches(db *gorm.DB) (*[]Batch, error) {
	var err error
	batches := []Batch{}
	err = db.Debug().Model(&User{}).Find(&batches).Error
	if err != nil {
		return &[]Batch{}, err
	}
	return &batches, err
}

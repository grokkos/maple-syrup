package models

import (
	"github.com/jinzhu/gorm"
)

type Batch struct {
	ID          uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Dispatched  bool   `gorm:"default:false"`
	Summary     int    `gorm:"not null" json:"summary"`
	BatchUserID uint32 `sql:"type:int REFERENCES users(id)" json:"batch_user_id"`
	BatchUser   User   `json:"batch_user"`
}
type Filter struct {
	ID int64 `schema:"id"`
}

var batch Batch

func (u *Batch) FindAllBatches(db *gorm.DB) (*[]Batch, error) {
	var err error
	batches := []Batch{}
	err = db.Debug().Model(&Batch{}).Find(&batches).Error
	if err != nil {
		return &[]Batch{}, err
	}
	return &batches, err
}

//getting all the batches filtered by user id
func (u *Batch) FindBatchesByUserId(db *gorm.DB, uid uint32) (*[]Batch, error) {
	var err error
	batches := []Batch{}
	err = db.Debug().Model(User{}).Where("batch_user_id = ?", uid).Find(&batches).Error

	if err != nil {
		return &[]Batch{}, err
	}
	return &batches, err
}

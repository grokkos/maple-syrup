package models

import (
	_ "github.com/jinzhu/gorm"
)

type Roundup struct {
	ID             uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Amount         int    `gorm:"not null" json:"amount"`
	RoundupBatchID uint32 `sql:"type:int REFERENCES batches(id)" json:"roundup_batch_id"`
	RoundupUserID  uint32 `sql:"type:int REFERENCES users(id)" json:"roundup_user_id"`
	RoundupUser    User   `json:"roundup_user"`
}

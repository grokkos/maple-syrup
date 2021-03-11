package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
)

type Roundup struct {
	ID             uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Amount         int    `gorm:"not null" json:"amount"`
	RoundupBatchID uint32 `sql:"type:int REFERENCES batches(id)" json:"roundup_batch_id"`
	RoundupUserID  uint32 `sql:"type:int REFERENCES users(id)" json:"roundup_user_id"`
	RoundupUser    User   `json:"roundup_user"`
}

func (r *Roundup) SaveRoundup(db *gorm.DB) (*Roundup, error) {

	var batch Batch
	var transaction Transaction
	var err error
	err = db.Debug().Model(&Batch{}).Where("dispatched = ?", false).Take(&batch).Error
	r.RoundupBatchID = batch.ID

	test := Roundup{
		ID:             r.ID,
		Amount:         r.Amount,
		RoundupBatchID: r.RoundupBatchID,
		RoundupUserID:  r.RoundupUserID,
	}
	err = db.Debug().Create(&test).Error

	if err != nil {
		return &Roundup{}, err
	}

	test2 := Batch{
		BatchUserID: r.RoundupUserID,
	}

	rows, err := db.Model(&Roundup{}).Where("roundup_batch_id = ?", r.RoundupBatchID).Select("amount").Rows()
	defer rows.Close()

	var round Roundup
	m := 0
	for rows.Next() {
		db.ScanRows(rows, &round)
		m += round.Amount
	}
	db.Model(&batch).Where("id = ?", batch.ID).Update("batch_user_id", r.RoundupUserID)

	if m > 100 {
		db.Model(&batch).Where("id = ?", batch.ID).Update("summary", m)
		db.Model(&batch).Where("id = ?", batch.ID).Update("dispatched", true)
		db.Debug().Model(&batch).Create(&test2)

		test3 := Transaction{
			TransactionUserID: r.RoundupUserID,
			Amount:            batch.Summary,
		}
		db.Debug().Model(&transaction).Create(&test3)
	}
	return r, nil
}

func (r *Roundup) FindAllRoundups(db *gorm.DB) (*[]Roundup, error) {
	var err error
	roundups := []Roundup{}
	err = db.Debug().Model(&Roundup{}).Find(&roundups).Error
	if err != nil {
		return &[]Roundup{}, err
	}
	return &roundups, err
}

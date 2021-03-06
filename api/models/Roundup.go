package models

import (
	"github.com/jinzhu/gorm"
)

type Roundup struct {
	ID             uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Amount         int    `gorm:"not null" json:"amount"`
	RoundupBatchID uint32 `sql:"type:int REFERENCES batches(id)" json:"roundup_batch_id"`
	RoundupUserID  uint32 `sql:"type:int REFERENCES users(id)" json:"roundup_user_id"`
}

func (r *Roundup) SaveRoundup(db *gorm.DB) (*Roundup, error) {

	var batch Batch
	var err error
	err = db.Debug().Model(&Batch{}).Where("dispatched = ?", false).Take(&batch).Error
	r.RoundupBatchID = batch.ID //when creating the roundup fetch the existing batch id which is not dispatched yet

	roundup := Roundup{ //define the roundup to be saved with the fetched batch id
		ID:             r.ID,
		Amount:         r.Amount,
		RoundupBatchID: r.RoundupBatchID,
		RoundupUserID:  r.RoundupUserID,
	}
	err = db.Debug().Create(&roundup).Error

	if err != nil {
		return &Roundup{}, err
	}

	nextbatch := Batch{ //set the user id for the generated batch
		BatchUserID: r.RoundupUserID,
	}

	rows, err := db.Model(&Roundup{}).Where("roundup_batch_id = ?", r.RoundupBatchID).Select("amount").Rows() //calculate the summary by batch id
	defer rows.Close()

	sum := 0
	for rows.Next() {
		db.ScanRows(rows, &roundup)
		sum += roundup.Amount
	}
	db.Model(&batch).Where("id = ?", batch.ID).Update("batch_user_id", r.RoundupUserID) //fetching the correct user id to the batch

	if sum > 100 {
		db.Model(&batch).Where("id = ?", batch.ID).Update("summary", sum)     //storing the sum only when it exceeds the threshold
		db.Model(&batch).Where("id = ?", batch.ID).Update("dispatched", true) //update the dispatched to trues
		db.Debug().Model(&batch).Create(&nextbatch)

		nexttransaction := Transaction{ //generate the transaction with the summary of amounts that is dispatched and the batch id
			TransactionBatchID: batch.ID,
			Amount:             batch.Summary,
		}
		db.Debug().Model(&nexttransaction).Create(&nexttransaction)
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

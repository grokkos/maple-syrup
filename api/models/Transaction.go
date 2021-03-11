package models

import (
	"github.com/jinzhu/gorm"
)

type Transaction struct {
	ID                 uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Amount             int    `gorm:"not null" json:"amount"`
	TransactionBatchID uint32 `sql:"type:int REFERENCES batches(id)" json:"transaction_batch_id"`
}

func (t *Transaction) FindAllTransactions(db *gorm.DB) (*[]Transaction, error) {
	var err error
	transactions := []Transaction{}
	err = db.Debug().Model(&Transaction{}).Find(&transactions).Error
	if err != nil {
		return &[]Transaction{}, err
	}
	return &transactions, err
}

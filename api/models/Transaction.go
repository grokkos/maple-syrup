package models

type Transaction struct {
	ID                 uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Amount             int    `gorm:"not null" json:"amount"`
	TransactionBatchID uint32 `sql:"type:int REFERENCES users(id)" json:"transaction_batch_id"`
}

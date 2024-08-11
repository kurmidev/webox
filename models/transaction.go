package models

import (
	"fmt"
	"time"

	"gorm.io/gorm/clause"
)

type Transaction struct {
	ID                             int `gorm:"primary_key"`
	SubsTranId                     int
	SchemeId                       int
	AccountId                      int
	SubscriberId                   int
	OperatorId                     int
	Amount                         float32
	Balance                        float32
	Tax                            float32
	Tds                            float32
	TdsOn                          float32
	Mrp                            float32
	TotalAmount                    float32
	Igst                           float32
	Cgst                           float32
	Sgst                           float32
	Type                           int
	Details                        string
	Remark                         string
	CreatedAt                      time.Time
	UpdatedAt                      time.Time
	CreatedBy                      int
	UpdatedBy                      int
	BouqueId                       int
	StartDate                      time.Time
	EndDate                        time.Time
	RecieptNo                      string
	PerDayAmount                   float32
	MrpTax                         float32
	PerDayMrp                      float32
	CancelledOperatorTransactionId int
	MrpBalance                     float32
	CMonth                         int
	IsRefundable                   int
	IsPartialRefund                int
	OrderId                        string
	Stbno                          string
	CustomerId                     string
	Smartcardno                    string
	IsCredit                       int
	Agr                            float32
	User                           User `gorm:"foreignKey:CreatedBy"`
}

func (t *Transaction) TableName() string {
	return "operator_transaction"
}

func (t *Transaction) GetTransactions(id int) ([]Transaction, error) {
	var transaction []Transaction
	sql := DB.Table(t.TableName()).
		Preload("User").
		Preload(clause.Associations).Where("operator_transaction.account_id=?", id)

	err := sql.Find(&transaction).Error
	fmt.Println("errr of the query", err)
	return transaction, err
}

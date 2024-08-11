package models

import "time"

type OperatorBalance struct {
	ID           int `gorm:"primary_key"`
	TotalCredit  float64
	TotalDebit   float64
	CreditCount  float64
	DebitCount   float64
	Balance      float64
	UpdatedAt    time.Time
	TotalCreditH float64
	TotalDebitH  float64
	CreditCountH float64
	DebitCountH  float64
	BalanceH     float64
	LastAccId    float64
}

func (opb *OperatorBalance) Table() string {
	return "operator_balance"
}

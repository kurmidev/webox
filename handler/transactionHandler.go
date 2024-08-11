package handler

type Transaction struct {
	ID                             int     `json:"id"`
	SubTranId                      int     `json:"subs_tran_id"`
	ReceiptNo                      string  `json:"reciept_no"`
	BouqueId                       int     `json:"bouque_id"`
	AccountId                      int     `json:"account_id"`
	SubscriberId                   int     `json:"subscriber_id"`
	OperatorId                     int     `json:"operator_id"`
	Amount                         float32 `json:"amount"`
	Balance                        float32 `json:"balance"`
	Type                           int     `json:"type"`
	Tds                            float32 `json:"tds"`
	Mrp                            float32 `json:"mrp"`
	TdsOn                          float32 `json:"tds_on"`
	StartDate                      string  `json:"start_date"`
	EndDate                        string  `json:"end_date"`
	CreatedAt                      string  `json:"created_at"`
	UpdatedAt                      string  `json:"updated_at"`
	CreatedBy                      int     `json:"created_by"`
	UpdatedBy                      int     `json:"updated_by"`
	SchemeId                       int     `json:"scheme_id"`
	Tax                            float32 `json:"tax"`
	TotalAmount                    float32 `json:"total_amount"`
	Igst                           float32 `json:"igst"`
	Cgst                           float32 `json:"cgst"`
	Sgst                           float32 `json:"sgst"`
	MrpTax                         float32 `json:"mrp_tax"`
	PerDayMrp                      float32 `json:"per_day_mrp"`
	Remark                         string  `json:"remark"`
	CancelledOperatorTransactionId int     `json:"cancelled_operator_transaction_id"`
	OrderId                        int     `json:"order_id"`
	CreatedByLbl                   string  `json:"created_by_lbl"`
	TypeLbl                        string  `json:"type_lbl"`
	NoteLbl                        string  `json:"note_lbl"`
}

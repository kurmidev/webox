package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

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
	OrderId                        string  `json:"order_id"`
	CreatedByLbl                   string  `json:"created_by_lbl"`
	TypeLbl                        string  `json:"type_lbl"`
	NoteLbl                        string  `json:"note_lbl"`
}

func (h *Handlers) ToTransactionResponse(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	transactions, err := h.Models.GetTransactions(id)
	if err != nil {
		fmt.Println(err)
		message := map[string]string{"number": "No transactions found"}
		h.Common.WriteJSON(w, http.StatusBadRequest, message)
		return
	}
	var response []Transaction
	for _, tx := range transactions {
		response = append(response, Transaction{
			ID:                             tx.ID,
			SubTranId:                      tx.SubsTranId,
			ReceiptNo:                      tx.RecieptNo,
			BouqueId:                       tx.BouqueId,
			AccountId:                      tx.AccountId,
			SubscriberId:                   tx.SubscriberId,
			OperatorId:                     tx.OperatorId,
			Amount:                         tx.Amount,
			Balance:                        tx.Balance,
			Type:                           tx.Type,
			Tds:                            tx.Tds,
			Mrp:                            tx.Mrp,
			TdsOn:                          tx.TdsOn,
			StartDate:                      tx.StartDate.Format("2006-01-02"),
			EndDate:                        tx.EndDate.Format("2006-01-02"),
			CreatedAt:                      tx.CreatedAt.Format("2006-01-02 15:04"),
			UpdatedAt:                      tx.UpdatedAt.Format("2006-01-02 15:04"),
			CreatedBy:                      tx.CreatedBy,
			UpdatedBy:                      tx.UpdatedBy,
			SchemeId:                       tx.SchemeId,
			Tax:                            tx.Tax,
			TotalAmount:                    tx.TotalAmount,
			Igst:                           tx.Igst,
			Cgst:                           tx.Cgst,
			Sgst:                           tx.Sgst,
			MrpTax:                         tx.MrpTax,
			PerDayMrp:                      tx.PerDayMrp,
			Remark:                         tx.Remark,
			CancelledOperatorTransactionId: tx.CancelledOperatorTransactionId,
			OrderId:                        tx.OrderId,
			CreatedByLbl:                   tx.User.Name,
		})
	}

	h.Common.WriteJSON(w, http.StatusOK, response)
}

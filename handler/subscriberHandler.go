package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/kurmidev/webox/utils"
)

type VcDetails struct {
	Cas          string `json:"cas"`
	OperatorName string `json:"operator_name"`
	OperatorCode string `json:"operator_code"`
	OperatorId   int    `json:"operator_id"`
	Smartcardno  string `json:"smartcardno"`
	Stbno        string `json:"stbno"`
	Brand        string `json:"brand"`
	IsHd         int    `json:"is_hd"`
	InvState     string `json:"inv_state"`
	OtherId      string `json:"other_id"`
	AccountId    int    `json:"account_id"`
}

type AccountDetails struct {
	CustomerId      string  `json:"customer_id"`
	Name            string  `json:"name"`
	MobileNo        string  `json:"mobile_no"`
	Address         string  `json:"address"`
	OperatorName    string  `json:"operator_name"`
	OperatorCode    string  `json:"operator_code"`
	OperatorId      int     `json:"operator_id"`
	OperatorBalance float32 `json:"operator_balance"`
	SubscriberId    int     `json:"subscriber_id"`
	AccountId       int     `json:"account_id"`
	Type            string  `json:"type"`
	Location        string  `json:"location"`
	SubLocation     string  `json:"sub_location"`
	OperatorCity    string  `json:"operator_city"`
	IsVerified      int     `json:"is_verified"`
}

type BouqueDetails struct {
	BouquetId        int     `json:"bouque_id"`
	BouquetTypeId    int     `json:"bouque_type_id"`
	BouquetType      string  `json:"bouque_type"`
	ActivationDate   string  `json:"activation_date"`
	DeactivationDate string  `json:"deactivation_date"`
	Amount           float32 `json:"amount"`
	Tax              float32 `json:"tax"`
	Left             int     `json:"left"`
	BouquetName      string  `json:"bouque_name"`
	Status           string  `json:"status"`
}

type Response struct {
	VcDetails  VcDetails                  `json:"vc_details"`
	AccDetails AccountDetails             `json:"account"`
	Connection map[string][]BouqueDetails `json:"connection"`
}

func (h *Handlers) SmcDetails(w http.ResponseWriter, r *http.Request) {
	number := chi.URLParam(r, "number")
	types, err := strconv.Atoi(chi.URLParam(r, "type"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	accdet, err := h.Models.GetAccount(number, types)
	if err != nil {
		fmt.Println(err)
		message := map[string]string{"number": "Invalid smc/stb details."}
		h.Common.WriteJSON(w, http.StatusBadRequest, message)
		return
	}

	vcDetails := VcDetails{
		Cas:          accdet.Cas.Name,
		OperatorName: accdet.Operator.Name,
		OperatorCode: accdet.Operator.Code,
		OperatorId:   accdet.OperatorId,
		Smartcardno:  accdet.Smartcardno,
		Stbno:        accdet.Stbno,
		Brand:        accdet.Brand.Name,
		IsHd:         accdet.IsHd,
		InvState:     "",
		OtherId:      "",
		AccountId:    accdet.ID,
	}

	accDetails := AccountDetails{
		CustomerId:      accdet.CustomerId,
		Name:            fmt.Sprintf("%s %s %s", accdet.Subscriber.Fname, accdet.Subscriber.Mname, accdet.Subscriber.Lname),
		MobileNo:        accdet.Subscriber.MobileNo,
		Address:         accdet.Subscriber.BillingAddress,
		OperatorName:    accdet.Operator.Name,
		OperatorCode:    accdet.Operator.Code,
		OperatorId:      accdet.OperatorId,
		OperatorBalance: float32(accdet.Operator.Balance.Balance),
		SubscriberId:    accdet.SubscriberId,
		AccountId:       accdet.ID,
		Type:            utils.GetCustomerType(accdet.Subscriber.CustomerType),
		Location:        accdet.Location.Name,
		SubLocation:     accdet.Sublocation.Name,
		OperatorCity:    accdet.Operator.City.Name,
		IsVerified:      accdet.Subscriber.IsVerified,
	}

	var bouquets []BouqueDetails

	for _, b := range accdet.SubscriberBouque {
		t1 := time.Date(b.ActivationDate.Year(), b.ActivationDate.Month(), b.ActivationDate.Day(), 0, 0, 0, 0, time.Local)
		t2 := time.Date(b.DeactivationDate.Year(), b.DeactivationDate.Month(), b.DeactivationDate.Day(), 0, 0, 0, 0, time.Local)
		days := t2.Sub(t1).Hours() / 24

		bouquets = append(bouquets, BouqueDetails{
			BouquetId:        b.BouqueId,
			BouquetTypeId:    b.BouqueType,
			BouquetType:      utils.BouqueTypeslbl(b.BouqueType),
			ActivationDate:   b.ActivationDate.Format("2006-12-01"),
			DeactivationDate: b.DeactivationDate.Format("2006-12-01"),
			Amount:           float32(b.Mrp),
			Tax:              float32(b.MrpTax),
			Left:             int(days),
			BouquetName:      b.Bouque.Name,
			Status:           utils.GetSubscriberBoquueStatus(b.Status),
		})
	}

	if len(bouquets) == 0 {
		for _, b := range accdet.InActiveSubscriberBouque {
			t1 := time.Date(b.ActivationDate.Year(), b.ActivationDate.Month(), b.ActivationDate.Day(), 0, 0, 0, 0, time.Local)
			t2 := time.Date(b.DeactivationDate.Year(), b.DeactivationDate.Month(), b.DeactivationDate.Day(), 0, 0, 0, 0, time.Local)
			days := t2.Sub(t1).Hours() / 24

			bouquets = append(bouquets, BouqueDetails{
				BouquetId:        b.BouqueId,
				BouquetTypeId:    b.BouqueType,
				BouquetType:      utils.BouqueTypeslbl(b.BouqueType),
				ActivationDate:   b.ActivationDate.Format("2006-12-01"),
				DeactivationDate: b.DeactivationDate.Format("2006-12-01"),
				Amount:           float32(b.Mrp),
				Tax:              float32(b.MrpTax),
				Left:             int(days),
				BouquetName:      b.Bouque.Name,
				Status:           utils.GetSubscriberBoquueStatus(b.Status),
			})
		}
	}

	response := Response{
		VcDetails:  vcDetails,
		AccDetails: accDetails,
		Connection: map[string][]BouqueDetails{
			"bouque": bouquets,
		},
	}

	h.Common.WriteJSON(w, http.StatusOK, response)
}

func (h *Handlers) ProfileDetails(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) MeDetails(w http.ResponseWriter, r *http.Request) {
}

/*
"vc_details": {
            "cas_lbl": "NSTV",
            "cas_code": "NSTV",
            "operator_name": "VD122L25-SWATHI COMMUNICATIONS",
            "operator_code": "VD122L25",
            "operator_id": 914,
            "cas_type": "NSTV",
            "sc_no": "000083330033D339",
            "sc_brand": "OVTVC",
            "stbno": "100309243669322",
            "stb_brand": "CHA045170A",
            "isHd": 1,
            "inv_state": "Alloted",
            "other_id": "N/A1",
            "account_id": "393773"
        },
        "account": {
            "customer_id": "UNI00427510",
            "lco_formno": null,
            "name": "custumer custumer",
            "mobile_no": "9999999999",
            "address": "meduru",
            "operator_name": "VD122L25-SWATHI COMMUNICATIONS",
            "operator_code": "VD122L25",
            "operator_id": 914,
            "operator_balance": "1757.4255",
            "subscriber_id": "427268",
            "account_id": "393773",
            "type": "Residential",
            "location": "VD122L25-SWATHI COMMUNICATIONS",
            "sub_location": "L25 SWATHI",
            "operator_city": "Gampalagudem mandal",
            "is_verified_lbl": "Not Verified"
        },
        "connection": {
            "bouque": [
                {
                    "bouque_id": 269,
                    "account_id": "393773",
                    "bouque_type": 1,
                    "activation_date": "2024-08-03",
                    "deactivation_date": "2024-09-02",
                    "amount": "0.0000",
                    "tax": "0.0000",
                    "left": "1 Month",
                    "bouque_name": "FOUNDATION+TELUGU ROYAL PROMO PACK",
                    "refundAble": {
                        "amt": 0,
                        "tax": 0,
                        "tds": 0,
                        "mrp": 357.25,
                        "mrp_tax": 0,
                        "non_refundable": 0
                    },
                    "status": 1,
                    "status_lbl": "Active"
                }
            ]
        }
    }*/

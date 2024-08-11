package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/kurmidev/webox/utils"
)

type SublocationDetail struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Code            string `json:"code"`
	Status          string `json:"status"`
	LocationId      int    `json:"location_id"`
	LocationLbl     string `json:"location_lbl"`
	OperatorId      int    `json:"operator_id"`
	Address         string `json:"address"`
	Email           string `json:"email"`
	MobileNo        string `json:"mobile_no"`
	OperatorName    string `json:"operator_lbl"`
	MsoId           int    `json:"mso_id"`
	Mso             string `json:"mso_lbl"`
	BranchId        int    `json:"branch_id"`
	BranchName      string `json:"branch_lbl"`
	DistributorId   int    `json:"distributor_id"`
	DistributorName string `json:"distributor_lbl"`
	OperatorCode    string `json:"operator_code_lbl"`
	City            string `json:"city_lbl"`
	DistrictId      int    `json:"district_lbl"`
	DistributorCode string `json:"distributor_code_lbl"`
	BranchCode      string `json:"branch_code_lbl"`
}

type Profile struct {
	ID                int               `json:"id"`
	CustomerId        string            `json:"customer_id"`
	Formno            string            `json:"formno"`
	Name              string            `json:"name"`
	Fname             string            `json:"fname"`
	Mname             string            `json:"lname"`
	Lname             string            `json:"mname"`
	BillingAddress    map[string]string `json:"billing_address"`
	Pincode           string            `json:"pincode"`
	Dob               string            `json:"dob"`
	MobileNo          string            `json:"mobile_no"`
	PhoneNo           string            `json:"phone_no"`
	Email             string            `json:"email"`
	Gender            string            `json:"gender"`
	CustomerType      string            `json:"customer_type"`
	Operator          string            `json:"operator"`
	OperatorCode      string            `json:"operator_code"`
	OperatorContactno string            `json:"operator_contactno"`
	SublocationDetail SublocationDetail `json:"sublocation_details"`
	EmailVerified     string            `json:"email_verified"`
	MobileNoVerified  string            `json:"mobile_no_verified"`
	Address           map[string]string `json:"address"`
}

type AcctBouque struct {
	BouqueId          int     `json:"bouque_id"`
	AccountId         int     `json:"account_id"`
	BouqueType        string  `json:"bouque_type"`
	ActivationDate    string  `json:"activation_date"`
	DeactivateionDate string  `json:"deactivation_date"`
	Amount            float32 `json:"amount"`
	Tax               float32 `json:"tax"`
	Left              int     `json:"left"`
	BouqueName        string  `json:"bouque_name"`
	Status            int     `json:"status"`
	StatusLbl         string  `json:"status_lbl"`
}

type AccDetails struct {
	ID               int          `json:"id"`
	Smartcardno      string       `json:"smartcardno"`
	Stbno            string       `json:"stbno"`
	ActivationDate   string       `json:"activation_date"`
	DeactivationDate string       `json:"deactivation_date"`
	Outstanding      float32      `json:"outstanding"`
	Status           string       `json:"status_lbl"`
	StatusInt        int          `json:"status"`
	BrandId          int          `json:"brand_id"`
	AcctBouque       []AcctBouque `json:"bouque"`
}

type ProfileDetails struct {
	Profile  Profile      `json:"profile"`
	Accounts []AccDetails `json:"accounts"`
}

func (h *Handlers) GetProfile(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	subscriber, err := h.Models.GetSubscriber(id)
	if subscriber.ID <= 0 {
		fmt.Println(err)
		message := map[string]string{"number": "Invalid profile id."}
		h.Common.WriteJSON(w, http.StatusBadRequest, message)
		return
	}

	var response ProfileDetails
	male := "Male"
	if subscriber.Gender == 0 {
		male = "Female"
	}
	response.Profile = Profile{
		ID:                subscriber.ID,
		CustomerId:        subscriber.CustomerId,
		Formno:            subscriber.Formno,
		Name:              fmt.Sprintf("%s %s %s", subscriber.Fname, subscriber.Mname, subscriber.Lname),
		Fname:             subscriber.Fname,
		Mname:             subscriber.Mname,
		Lname:             subscriber.Lname,
		BillingAddress:    utils.FormatJson(subscriber.BillingAddress),
		Pincode:           subscriber.Pincode,
		Dob:               subscriber.Dob.Format("2006-01-02"),
		MobileNo:          subscriber.MobileNo,
		PhoneNo:           subscriber.PhoneNo,
		Email:             subscriber.Email,
		Gender:            male,
		CustomerType:      utils.GetCustomerType(subscriber.CustomerType),
		Operator:          subscriber.Operator.Name,
		OperatorCode:      subscriber.Operator.Code,
		OperatorContactno: subscriber.Operator.MobileNo,
		SublocationDetail: SublocationDetail{
			ID:              subscriber.SublocationId,
			Name:            subscriber.Sublocation.Name,
			Code:            subscriber.Sublocation.Code,
			LocationId:      subscriber.LocationId,
			LocationLbl:     subscriber.Location.Name,
			OperatorId:      subscriber.OperatorId,
			Address:         subscriber.Operator.Addr,
			Email:           subscriber.Operator.Email,
			MobileNo:        subscriber.Operator.MobileNo,
			OperatorName:    subscriber.Operator.Name,
			MsoId:           subscriber.Operator.MsoId,
			Mso:             subscriber.Operator.Mso.Name,
			BranchId:        subscriber.Operator.BranchId,
			BranchName:      subscriber.Operator.Branch.Name,
			BranchCode:      subscriber.Operator.Branch.Code,
			DistributorId:   subscriber.Operator.DistributorId,
			DistributorName: subscriber.Operator.Distributor.Name,
			DistributorCode: subscriber.Operator.Distributor.Code,
			OperatorCode:    subscriber.Operator.Code,
			City:            subscriber.Operator.City.Name,
			DistrictId:      subscriber.Operator.DistrictId,
		},
		EmailVerified:    utils.IsApp(subscriber.IsVerified),
		MobileNoVerified: utils.IsApp(subscriber.IsVerified),
		Address:          utils.FormatJson(subscriber.InstallationAddress),
	}
	var acc []AccDetails
	for _, ac := range subscriber.SubscriberAccount {
		var accbous []AcctBouque

		for _, accbou := range ac.SubscriberBouque {
			t1 := time.Date(accbou.ActivationDate.Year(), accbou.ActivationDate.Month(), accbou.ActivationDate.Day(), 0, 0, 0, 0, time.Local)
			t2 := time.Date(accbou.DeactivationDate.Year(), accbou.DeactivationDate.Month(), accbou.DeactivationDate.Day(), 0, 0, 0, 0, time.Local)
			days := t2.Sub(t1).Hours() / 24

			accbous = append(accbous, AcctBouque{
				BouqueId:          accbou.BouqueId,
				AccountId:         accbou.AccountId,
				BouqueType:        utils.BouqueTypeslbl(accbou.BouqueType),
				ActivationDate:    accbou.ActivationDate.Format("2006-01-06"),
				DeactivateionDate: accbou.DeactivationDate.Format("2006-01-06"),
				Amount:            float32(accbou.Mrp),
				Tax:               float32(accbou.MrpTax),
				Left:              int(days),
				BouqueName:        accbou.Bouque.Name,
				Status:            accbou.Status,
				StatusLbl:         utils.GetSubscriberBoquueStatus(accbou.Status),
			})
		}

		acc = append(acc, AccDetails{
			ID:               ac.ID,
			Smartcardno:      ac.Smartcardno,
			Stbno:            ac.Stbno,
			ActivationDate:   ac.ActivationDate.Format("2006-01-06"),
			DeactivationDate: ac.DeactivationDate.Format("2006-01-06"),
			Outstanding:      0,
			Status:           utils.GetSubscriberBoquueStatus(ac.Status),
			StatusInt:        ac.Status,
			BrandId:          ac.StbbrandId,
			AcctBouque:       accbous,
		})
	}

	response.Accounts = acc

	h.Common.WriteJSON(w, http.StatusOK, response)
}

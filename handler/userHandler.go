package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kurmidev/webox/utils"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	LoginForm `json:"LoginForm"`
}

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Id          int    `json:"id"`
	Role        int    `json:"role"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	MobileNo    string `json:"mobile_no"`
	AccessToken string `json:"access_token"`
	AuthToken   string `json:"auth_token"`
	Expiry      string `json:"expiry"`
}

type LoginOtpRequest struct {
	MobileNo string `json:"mobile_no"`
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var logindata LoginRequest
	var payload LoginResponse
	err := json.NewDecoder(r.Body).Decode(&logindata)
	if err != nil {
		message := map[string]string{"request": "Invalid request format."}
		_ = h.Common.WriteJSON(w, http.StatusBadRequest, message)
		return
	}

	user, err := h.Models.User.GetUser(logindata.Username)
	if err != nil {
		message := map[string]string{"username": "Invalid username details provided."}
		h.Common.WriteJSON(w, http.StatusBadRequest, message)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(logindata.LoginForm.Password))
	if err != nil {
		message := map[string]string{"password": "Invalid password details provided."}
		h.Common.WriteJSON(w, http.StatusBadRequest, message)
		return
	}

	var currentTime = time.Now().Unix()
	var expire = time.Now().Add(100000 * time.Minute).Unix()

	claim := jwt.MapClaims{}
	claim["iat"] = currentTime
	claim["iss"] = fmt.Sprintf("http://%s", r.Host)
	claim["aud"] = fmt.Sprintf("http://%s", r.Host)
	claim["nbf"] = currentTime
	claim["exp"] = expire
	claim["is_aggrement_void"] = 0
	claim["allowed_apis"] = nil
	claim["extra_data"] = []string{}
	claim["restrict_ip"] = []string{}
	claim["jti"] = user.Id
	claim["data"] = map[string]interface{}{
		"username":    user.Username,
		"roleLabel":   utils.GetRoles(user.Role),
		"lastLoginAt": user.LastLoginAt.Format("2006-01-02 15:04:05"),
		"session_id":  "",
		"auth_key":    user.AuthKey,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	payload.AccessToken, _ = token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	payload.Username = user.Username
	payload.AuthToken = user.AuthKey
	payload.MobileNo = user.MobileNo
	payload.Email = user.Email
	payload.Id = user.Id
	payload.Name = user.Name
	payload.Role = user.Role
	payload.Expiry = time.Now().Add(100000 * time.Minute).Format("2006-01-02 15:04:05")
	h.Common.WriteJSON(w, http.StatusOK, payload)

}

func (h *Handlers) SendOtp(w http.ResponseWriter, r *http.Request) {
	var logindata LoginOtpRequest

	err := json.NewDecoder(r.Body).Decode(&logindata)
	if err != nil {
		message := map[string]string{"request": "Invalid request format."}
		_ = h.Common.WriteJSON(w, http.StatusBadRequest, message)
		return
	}

	user,err;= h.Models.GetUserByMobile(lologindata.MobileNo)
	if err!= nil {
        message := map[string]string{"mobile_no": "Invalid mobile number details provided."}
        h.Common.WriteJSON(w, http.StatusBadRequest, message)
        return
    }
}

func (h *Handlers) LoginOtp(w http.ResponseWriter, r *http.Request) {
	var logindata LoginRequest
	var payload LoginResponse
	err := json.NewDecoder(r.Body).Decode(&logindata)
	if err != nil {
		message := map[string]string{"request": "Invalid request format."}
		_ = h.Common.WriteJSON(w, http.StatusBadRequest, message)
		return
	}

	user, err := h.Models.User.GetUser(logindata.Username)
	if err != nil {
		message := map[string]string{"username": "Invalid username details provided."}
		h.Common.WriteJSON(w, http.StatusBadRequest, message)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(logindata.LoginForm.Password))
	if err != nil {
		message := map[string]string{"password": "Invalid password details provided."}
		h.Common.WriteJSON(w, http.StatusBadRequest, message)
		return
	}

	var currentTime = time.Now().Unix()
	var expire = time.Now().Add(100000 * time.Minute).Unix()

	claim := jwt.MapClaims{}
	claim["iat"] = currentTime
	claim["iss"] = fmt.Sprintf("http://%s", r.Host)
	claim["aud"] = fmt.Sprintf("http://%s", r.Host)
	claim["nbf"] = currentTime
	claim["exp"] = expire
	claim["is_aggrement_void"] = 0
	claim["allowed_apis"] = nil
	claim["extra_data"] = []string{}
	claim["restrict_ip"] = []string{}
	claim["jti"] = user.Id
	claim["data"] = map[string]interface{}{
		"username":    user.Username,
		"roleLabel":   utils.GetRoles(user.Role),
		"lastLoginAt": user.LastLoginAt.Format("2006-01-02 15:04:05"),
		"session_id":  "",
		"auth_key":    user.AuthKey,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	payload.AccessToken, _ = token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	payload.Username = user.Username
	payload.AuthToken = user.AuthKey
	payload.MobileNo = user.MobileNo
	payload.Email = user.Email
	payload.Id = user.Id
	payload.Name = user.Name
	payload.Role = user.Role
	payload.Expiry = time.Now().Add(100000 * time.Minute).Format("2006-01-02 15:04:05")
	h.Common.WriteJSON(w, http.StatusOK, payload)
}

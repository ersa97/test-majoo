package testmajoo

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ersa97/test-majoo/helpers"
	"github.com/ersa97/test-majoo/models"
	"github.com/jinzhu/gorm"
)

type MajooService struct {
	DB *gorm.DB
}

func (m *MajooService) TestAPI(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "API is working",
		Data:    nil,
	})
}

func (m *MajooService) Login(w http.ResponseWriter, r *http.Request) {

	var body models.User
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "Get Body Failed",
		})
		return
	}

	result, err := models.GetUser(body.Username, m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}

	if fmt.Sprintf("%x", md5.Sum([]byte(body.Password))) != result.Password {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "wrong password",
		})
		return
	}
	token := helpers.GenerateToken(result.Id, result.Name, body.Username)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "login success",
		Data: map[string]interface{}{
			"token": token,
		},
	})
}

func (m *MajooService) GetMerchantOmzet(w http.ResponseWriter, r *http.Request) {

	id := int(helpers.GetAuthorizationTokenValue(r, "userid").(float64))

	limit, err := strconv.Atoi(r.URL.Query()["limit"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "required field limit",
		})
		return
	}
	page, err := strconv.Atoi(r.URL.Query()["page"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "required field limit",
		})
		return
	}

	result, err := models.GetMerchantOmzetByUserId(id, limit, page, m.DB)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "get merchant omzet success",
		Data:    result,
	})
}

func (m *MajooService) GetMerchantOutletOmzet(w http.ResponseWriter, r *http.Request) {
	id := int(helpers.GetAuthorizationTokenValue(r, "userid").(float64))

	limit, err := strconv.Atoi(r.URL.Query()["limit"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "required field limit",
		})
		return
	}
	page, err := strconv.Atoi(r.URL.Query()["page"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Message: "required field limit",
		})
		return
	}

	result, err := models.GetMerchantOutletOmzetByUserId(id, limit, page, m.DB)
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Response{
		Message: "get merchant outlet omzet success",
		Data:    result,
	})
}

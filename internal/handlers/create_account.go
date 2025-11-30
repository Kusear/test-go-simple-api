package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"test-go-simple-api/api"
	"test-go-simple-api/internal/database"
)

type NewAccountResponse struct {
	Username string
	Balance  int64
	ID       int
}

func PostCreateAccount(w http.ResponseWriter, r *http.Request) {
	// username := r.URL.Query().Get("username")
	// balance := getCoinBalance(username)
	// w.WriteHeader(http.StatusOK)

	db := database.GetDatabaseInstance()

	account := NewAccountResponse{
		Username: "test",
		Balance:  228,
	}

	item, err := db.CreateItem(account)
	if err != nil {
		api.InternalServerError(w)
		return
	}

	accountData, ok := item.Data.(NewAccountResponse)
	if !ok {
		api.InternalServerError(w)
		return
	}

	w.Header().Add("X-Auth-Token", strings.Join([]string{fmt.Sprintf("%d", item.ID), accountData.Username}, ":"))
	api.SuccessResponse(w, api.ResponseData[NewAccountResponse]{
		Code: http.StatusOK,
		Data: NewAccountResponse{
			Username: accountData.Username,
			Balance:  accountData.Balance,
			ID:       item.ID,
		},
	})
}

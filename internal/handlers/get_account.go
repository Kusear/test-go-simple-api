package handlers

import (
	"net/http"

	"test-go-simple-api/api"
	"test-go-simple-api/internal/database"
)

func GetAccount(w http.ResponseWriter, r *http.Request) {
	accountId := r.Context().Value("accountId")

	db := database.GetDatabaseInstance()

	account, err := db.GetItem(accountId.(int))
	if err != nil {
		api.UnauthorizedError(w)
		return
	}

	api.SuccessResponse(w, api.ResponseData[NewAccountResponse]{
		Code: http.StatusOK,
		Data: NewAccountResponse{
			Username: account.Data.(NewAccountResponse).Username,
			Balance:  account.Data.(NewAccountResponse).Balance,
			ID:       account.ID,
		},
	})
}

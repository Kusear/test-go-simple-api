package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"test-go-simple-api/api"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("X-Auth-Token")
		if authToken == "" {
			api.UnauthorizedError(w)
			return
		}

		tokenParts := strings.Split(authToken, ":")
		if len(tokenParts) != 2 {
			api.UnauthorizedError(w)
			return
		}

		accountID, err := strconv.Atoi(tokenParts[0])
		if err != nil {
			api.UnauthorizedError(w)
			return
		}

		// TODO account existing check
		// db := database.GetDatabaseInstance()
		// _, err = db.GetItem(accountID)
		// if err != nil {
		// 	api.UnauthorizedError(w)
		// 	return
		// }

		ctx := context.WithValue(r.Context(), "accountId", accountID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

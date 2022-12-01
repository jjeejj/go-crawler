// middleware WithAuth 判断是否登录
package middleware

import (
	"context"
	"net/http"
)

func WithAuth() http.Handler {

	context.Background()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	})
}

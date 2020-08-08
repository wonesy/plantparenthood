package auth

import (
	"context"
	"net/http"

	"github.com/wonesy/plantparenthood/pkg/jwt"
)

var memberCtxKey = &contextKey{"id"}

type contextKey struct {
	id string
}

// CheckTokenMiddleware middleware for authentication of jwt
func CheckTokenMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			id, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), memberCtxKey, &id)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// IDFromContext finds the member ID from the context. REQUIRES Middleware to have run.
func IDFromContext(ctx context.Context) string {
	raw, _ := ctx.Value(memberCtxKey).(string)
	return raw
}

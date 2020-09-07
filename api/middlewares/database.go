package middlewares

import (
	"context"
	"database/sql"
	"net/http"
)

const (
	dbContext = "database"
)

func Database(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, dbContext, db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetDB(ctx context.Context) *sql.DB {
	return ctx.Value(dbContext).(*sql.DB)
}

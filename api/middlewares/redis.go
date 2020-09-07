package middlewares

import (
	"context"
	"net/http"

	"github.com/go-redis/redis"
)

const (
	redisContext = "redis"
)

func Redis(rds *redis.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, redisContext, rds)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetRedis(ctx context.Context) *redis.Client {
	return ctx.Value(redisContext).(*redis.Client)
}

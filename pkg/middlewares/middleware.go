package middlewares

import (
	"context"

	"github.com/dimovnike/forwardauth/pkg/log"
)

// GetLoggerCtx creates a logger context with the middleware fields.
func GetLoggerCtx(ctx context.Context, middleware, middlewareType string) context.Context {
	return log.With(ctx, log.Str(log.MiddlewareName, middleware), log.Str(log.MiddlewareType, middlewareType))
}

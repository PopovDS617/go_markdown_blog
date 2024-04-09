package middleware

import (
	"fmt"
	"gomarkdownblog/internal/utils"
	"net/http"

	"go.uber.org/zap"
)

func LoggingMiddleware(wrappedHandler http.HandlerFunc, logger *zap.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		lrw := utils.NewLoggingResponseWriter(w)
		wrappedHandler.ServeHTTP(lrw, r)

		statusCode := lrw.StatusCode

		logger.Info(fmt.Sprintf("%s %s %d", r.Method, r.URL.Path, statusCode))
	})
}

package middleware

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"noraclock/src/logger"
	"time"
)

var log = logger.General()

// Interceptor attaches the correlationID to the request context and logs its
// execution time. It is intended to be at the top of the middleware execution chain.
func Interceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		arrival := time.Now()
		reqCorrelationID := uuid.New().String()

		newCtx := context.WithValue(req.Context(), CorrelationIDKey, reqCorrelationID)
		req = req.WithContext(newCtx)

		next.ServeHTTP(writer, req)
		log.Sugar().Infof("middleware.Interceptor: Request took %dms to process.", time.Since(arrival).Milliseconds())
	})
}

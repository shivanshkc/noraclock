package middleware

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"noraclock/src/configs"
	"noraclock/src/exception"
	"noraclock/src/logger"
	"strings"
	"time"
)

var conf = configs.Get()
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

// CORS : Handles the CORS problems.
// TODO : Make the allowed origins configurable.
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "*")

		if req.Method == http.MethodOptions {
			writer.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(writer, req)
	})
}

// NoraGuard : Allows access to Nora user only.
func NoraGuard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		if !strings.HasPrefix(req.URL.String(), "/api/noraAccess") {
			next.ServeHTTP(writer, req)
			return
		}

		password := req.Header.Get("x-password")
		if password != conf.Service.Password {
			exception.Send(exception.Unauthorized(), writer)
			return
		}
		next.ServeHTTP(writer, req)
	})
}

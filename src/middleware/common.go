package middleware

import (
	"net/http"
	"noraclock/v2/src/configs"
	"noraclock/v2/src/exception"
	"noraclock/v2/src/logger"
	"strings"
	"time"
)

// Interceptor : The fist stage of request acceptance. Also acts as access logger.
func Interceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		arrival := time.Now()
		log, _ := logger.Get()

		log.Sugar().Infof("New Request %s %s, Timestamp: %d", req.Method, req.URL, arrival.UnixNano())
		next.ServeHTTP(writer, req)
		log.Sugar().Infof("Request took %dus to process.", time.Since(arrival).Microseconds())
	})
}

// CORS : Handles the CORS problems.
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "*")

		if req.Method == http.MethodOptions {
			writer.WriteHeader(200)
			return
		}
		next.ServeHTTP(writer, req)
	})
}

// NoraGuard : Allows access to Nora user only.
func NoraGuard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		password := req.Header.Get("x-password")
		if password != configs.Service.Password {
			exception.Unauthorized("").Send(writer)
			return
		}
		next.ServeHTTP(writer, req)
	})
}

// ResponseHeader : Sets common Response Headers.
func ResponseHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		if strings.HasPrefix(req.URL.String(), "/api/noraAccess") {
			writer.Header().Set("content-type", "application/json")
		}
		next.ServeHTTP(writer, req)
	})
}

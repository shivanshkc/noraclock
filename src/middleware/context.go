package middleware

// ReqContextKey is the type of the key that is
// used to put values in the HTTP request's context.
type ReqContextKey string

const (
	// CorrelationIDKey is used to put the request's correlation ID in its context.
	CorrelationIDKey ReqContextKey = "correlationID"
)

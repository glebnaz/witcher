package trace

import (
	"context"

	"github.com/google/uuid"
)

type simpleReqIDKey string

var simpleReqID simpleReqIDKey = "simpleReqID"

// AddSimpleReqIDToContext adds the simple request ID to the context
func AddSimpleReqIDToContext(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, simpleReqID, reqID)
}

// GetSimpleReqIDFromContext returns the simple request ID from the context
// if no simple request ID is found, returns an empty string
func GetSimpleReqIDFromContext(ctx context.Context) string {
	reqID, ok := ctx.Value(simpleReqID).(string)
	if !ok {
		return ""
	}
	return reqID
}

// MustGetSimpleReqIDFromContext returns the simple request ID from the context
// if no simple request ID is found, generates a new one and returns it
// and adds it to the context
// if simple request ID is found, returns new ctx with the simple request ID
func MustGetSimpleReqIDFromContext(ctx context.Context) (context.Context, string) {
	reqID, ok := ctx.Value(simpleReqID).(string)
	if !ok {
		reqID = GenerateSimpleReqID()
		ctx = AddSimpleReqIDToContext(ctx, reqID)
		return ctx, reqID
	}
	return ctx, reqID
}

// GenerateSimpleReqID generates a simple request ID
func GenerateSimpleReqID() string {
	return uuid.New().String()
}

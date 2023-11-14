package trace

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

type simpleReqIDKey string

var simpleReqID simpleReqIDKey = "simple-req-id"

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

// GetSimpleReqIDFromMetaData returns the simple request ID from the metadata
// metadata is incoming metadata
// if req-id contains in metadata we will add it to context
func GetSimpleReqIDFromMetaData(ctx context.Context) (string, context.Context) {
	incoming, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ctx
	}

	reqIDArray := incoming.Get(string(simpleReqID))

	if len(reqIDArray) > 0 {
		reqID := reqIDArray[0]
		ctx = context.WithValue(ctx, simpleReqID, reqID)
		return reqID, ctx
	}

	return "", ctx
}

// MustGetSimpleReqIDFromMetaData returns the simple request ID from the metadata
// metadata is incoming metadata
// if req-id not contains in metadata we will generate new one and add it to context
func MustGetSimpleReqIDFromMetaData(ctx context.Context) (string, context.Context) {
	returnFunc := func(ctx context.Context, reqID string) (string, context.Context) {
		ctx = context.WithValue(ctx, simpleReqID, reqID)
		return reqID, ctx
	}

	incoming, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return returnFunc(ctx, GenerateSimpleReqID())
	}

	reqIDArray := incoming.Get(string(simpleReqID))

	if len(reqIDArray) > 0 {
		reqID := reqIDArray[0]
		ctx = context.WithValue(ctx, simpleReqID, reqID)
		return reqID, ctx
	}

	return returnFunc(ctx, GenerateSimpleReqID())
}

// GenerateSimpleReqID generates a simple request ID
func GenerateSimpleReqID() string {
	return uuid.New().String()
}

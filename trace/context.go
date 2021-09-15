package trace

import (
	"context"

	"github.com/teris-io/shortid"
)

const (
	reqIDKey = "req-id"
	unknown  = "unknown"
)

//GenerateReqIDToCTX to ctx and return it
//if error of generate not nil return unknown
func GenerateReqIDToCTX(ctx context.Context) (context.Context, string) {
	ctx = AddNewReqIDToCTX(ctx, "")
	return ctx, GetReqIDFromCTX(ctx)
}

//AddNewReqIDToCTX add req id to ctx
//if req-id is empty generate new
func AddNewReqIDToCTX(ctx context.Context, reqID string) context.Context {
	if len(reqID) != 0 {
		return context.WithValue(ctx, reqIDKey, reqID)
	}
	return generateReqIDToCTX(ctx)
}

func generateReqIDToCTX(ctx context.Context) context.Context {
	reqID, err := shortid.Generate()
	if err != nil {
		reqID = unknown
	}

	return context.WithValue(ctx, reqIDKey, reqID)
}

//GetReqIDFromCTX return reqID
//if reqID in ctx empty get unknown
func GetReqIDFromCTX(ctx context.Context) string {
	reqID, ok := ctx.Value(reqIDKey).(string)
	if !ok || len(reqID) == 0 {
		return unknown
	}
	return reqID
}

package log

import (
	"github.com/glebnaz/witcher/trace"
	"github.com/rs/zerolog"
)

type fullInfoKeyType string

var fullInfoCtxKey = fullInfoKeyType("full-info")

// GRPCServerMWHook will add req-id and handler name to your log
type GRPCServerMWHook struct{}

func (l GRPCServerMWHook) Run(e *zerolog.Event, _ zerolog.Level, _ string) {
	ctx := e.GetCtx()
	reqID := trace.GetSimpleReqIDFromContext(ctx)
	if reqID != "" {
		e.Str("req-id", reqID)
	}
	fullInfo, ok := ctx.Value(fullInfoCtxKey).(string)
	if ok {
		e.Str("handler", fullInfo)
	}
}

var defaultHooks = []zerolog.Hook{
	GRPCServerMWHook{},
}

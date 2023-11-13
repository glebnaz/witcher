package log

import (
	"context"
	"log/slog"
)

type CTXAttrHandler struct {
	slog.Handler
}

// NewCTXAttrHandler
//
// is a slog.Handler that adds slog.Attr to slog.Record from context.Context
func NewCTXAttrHandler(serviceName string, mainHandler slog.Handler) CTXAttrHandler {
	mainHandler.WithAttrs([]slog.Attr{
		slog.String("service", serviceName),
	},
	)

	return CTXAttrHandler{
		Handler: mainHandler,
	}
}

func (p CTXAttrHandler) Handle(ctx context.Context, record slog.Record) error {
	attrs := GetSLOGAttrFromCTX(ctx)
	if attrs == nil {
		return p.Handler.Handle(ctx, record)
	}

	record.AddAttrs(attrs...)

	return p.Handler.Handle(ctx, record)
}

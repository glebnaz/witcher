package log

import (
	"context"
	"log/slog"
)

type ctxKeySLOGAttr string

var ctxKeySlOGAttr = ctxKeySLOGAttr("attr")

// AddSLOGAttrToCTX adds slog attr to ctx
func AddSLOGAttrToCTX(ctx context.Context, attrs []slog.Attr) context.Context {
	oldAttrs := GetSLOGAttrFromCTX(ctx)
	if oldAttrs != nil {
		attrs = append(oldAttrs, attrs...)
	}
	return context.WithValue(ctx, ctxKeySlOGAttr, attrs)
}

// GetSLOGAttrFromCTX returns slog attrs from ctx
func GetSLOGAttrFromCTX(ctx context.Context) []slog.Attr {
	attrs, ok := ctx.Value(ctxKeySlOGAttr).([]slog.Attr)
	if !ok {
		attrs = nil
	}
	return attrs
}

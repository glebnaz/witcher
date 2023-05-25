package log

import (
	"context"
	"github.com/sirupsen/logrus"
)

// for fields
type ctxFieldsKey string

var key = ctxFieldsKey("logrusFields")

// for logrus.Entry
type ctxEntryKey string

var entryKey = ctxEntryKey("logrusEntry")

// AddFieldsToCTX adds fields to ctx
func AddFieldsToCTX(ctx context.Context, fields logrus.Fields) context.Context {
	return context.WithValue(ctx, key, fields)
}

func GetFieldsFromCTX(ctx context.Context) logrus.Fields {
	fields, ok := ctx.Value(key).(logrus.Fields)
	if !ok {
		fields = logrus.Fields{}
	}
	return fields
}

// NewFromCTX returns a new logrus.Entry with fields from ctx
func NewFromCTX(ctx context.Context) *logrus.Entry {
	fields, ok := ctx.Value(key).(logrus.Fields)
	if !ok {
		fields = logrus.Fields{}
	}
	return logrus.WithFields(fields)
}

// AddEntryToCTX adds entry to ctx
func AddEntryToCTX(ctx context.Context, entry *logrus.Entry) context.Context {
	return context.WithValue(ctx, entryKey, entry)
}

// MustGetEntryFromCTX returns a *logrus.Entry from ctx
// if ctx has no *logrus.Entry, returns a new *logrus.Entry
func MustGetEntryFromCTX(ctx context.Context) *logrus.Entry {
	entry, ok := ctx.Value(entryKey).(*logrus.Entry)
	if !ok {
		return logrus.NewEntry(logrus.StandardLogger())
	}
	return entry
}

func GetEntryFromCTX(ctx context.Context) (*logrus.Entry, bool) {
	entry, ok := ctx.Value(entryKey).(*logrus.Entry)
	if !ok {
		return nil, false
	}
	return entry, true
}

// NewEntryFormCTXWithFields returns a *logrus.Entry with fields from ctx
func NewEntryFormCTXWithFields(ctx context.Context) *logrus.Entry {
	entry := MustGetEntryFromCTX(ctx)
	fields := GetFieldsFromCTX(ctx)
	return entry.WithFields(fields)
}

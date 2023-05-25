package log

import (
	"bytes"
	"context"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestInfof(t *testing.T) {
	bufferForCTX := &bytes.Buffer{}
	withCTXLogger := logrus.New()
	withCTXLogger.SetOutput(bufferForCTX)
	withCTXEntry := logrus.NewEntry(withCTXLogger)

	ctxWithEntry := AddEntryToCTX(context.Background(), withCTXEntry.WithFields(logrus.Fields{"test": "test"}))

	bufferWithoutCTX := &bytes.Buffer{}
	logrus.SetOutput(bufferWithoutCTX)

	type args struct {
		ctx    context.Context
		format string
		args   []interface{}
	}
	tests := []struct {
		name   string
		args   args
		buffer *bytes.Buffer
	}{
		{
			name: "with ctx",
			args: args{
				ctx:    ctxWithEntry,
				format: "test",
				args:   []interface{}{},
			},
			buffer: bufferForCTX,
		},
		{
			name: "without ctx",
			args: args{
				ctx:    context.Background(),
				format: "test",
				args:   []interface{}{},
			},
			buffer: bufferWithoutCTX,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Infof(tt.args.ctx, tt.args.format, tt.args.args...)
			if len(tt.buffer.String()) == 0 {
				t.Errorf("Infof() = %v, want %v", tt.buffer.String(), tt.args.format)
			}
		})
	}
}

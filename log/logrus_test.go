package log

import (
	"bytes"
	"context"
	"testing"

	"github.com/sirupsen/logrus"
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

func TestInfo(t *testing.T) {
	bufferForCTX := &bytes.Buffer{}
	withCTXLogger := logrus.New()
	withCTXLogger.SetOutput(bufferForCTX)
	withCTXEntry := logrus.NewEntry(withCTXLogger)

	ctxWithEntry := AddEntryToCTX(context.Background(), withCTXEntry.WithFields(logrus.Fields{"test": "test"}))

	bufferWithoutCTX := &bytes.Buffer{}
	logrus.SetOutput(bufferWithoutCTX)

	type args struct {
		ctx  context.Context
		args []interface{}
	}
	tests := []struct {
		name   string
		args   args
		buffer *bytes.Buffer
	}{
		{
			name: "with ctx",
			args: args{
				ctx:  ctxWithEntry,
				args: []interface{}{},
			},
			buffer: bufferForCTX,
		},
		{
			name: "without ctx",
			args: args{
				ctx:  context.Background(),
				args: []interface{}{},
			},
			buffer: bufferWithoutCTX,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.ctx, tt.args.args...)
			if len(tt.buffer.String()) == 0 {
				t.Errorf("Info() = %v, want %v", tt.buffer.String(), tt.args.args)
			}
		})
	}
}

func TestDebugf(t *testing.T) {
	bufferForCTX := &bytes.Buffer{}
	withCTXLogger := logrus.New()
	withCTXLogger.SetLevel(logrus.DebugLevel)
	withCTXLogger.SetOutput(bufferForCTX)
	withCTXEntry := logrus.NewEntry(withCTXLogger)

	ctxWithEntry := AddEntryToCTX(context.Background(), withCTXEntry.WithFields(logrus.Fields{"test": "test"}))

	bufferWithoutCTX := &bytes.Buffer{}
	logrus.SetLevel(logrus.DebugLevel)
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
			Debugf(tt.args.ctx, tt.args.format, tt.args.args...)
			if len(tt.buffer.String()) == 0 {
				t.Errorf("Debugf() = %v, want %v", tt.buffer.String(), tt.args.format)
			}
		})
	}
}

func TestDebug(t *testing.T) {
	bufferForCTX := &bytes.Buffer{}
	withCTXLogger := logrus.New()
	withCTXLogger.SetLevel(logrus.DebugLevel)
	withCTXLogger.SetOutput(bufferForCTX)
	withCTXEntry := logrus.NewEntry(withCTXLogger)

	ctxWithEntry := AddEntryToCTX(context.Background(), withCTXEntry.WithFields(logrus.Fields{"test": "test"}))

	bufferWithoutCTX := &bytes.Buffer{}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(bufferWithoutCTX)

	type args struct {
		ctx  context.Context
		args []interface{}
	}
	tests := []struct {
		name   string
		args   args
		buffer *bytes.Buffer
	}{
		{
			name: "with ctx",
			args: args{
				ctx:  ctxWithEntry,
				args: []interface{}{},
			},
			buffer: bufferForCTX,
		},
		{
			name: "without ctx",
			args: args{
				ctx:  context.Background(),
				args: []interface{}{},
			},
			buffer: bufferWithoutCTX,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.ctx, tt.args.args...)
			if len(tt.buffer.String()) == 0 {
				t.Errorf("Debug() = %v, want %v", tt.buffer.String(), tt.args.args)
			}
		})
	}
}

func TestErrorf(t *testing.T) {
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
			Errorf(tt.args.ctx, tt.args.format, tt.args.args...)
			if len(tt.buffer.String()) == 0 {
				t.Errorf("Errorf() = %v, want %v", tt.buffer.String(), tt.args.format)
			}
		})
	}
}

func TestError(t *testing.T) {
	bufferForCTX := &bytes.Buffer{}
	withCTXLogger := logrus.New()
	withCTXLogger.SetOutput(bufferForCTX)
	withCTXEntry := logrus.NewEntry(withCTXLogger)

	ctxWithEntry := AddEntryToCTX(context.Background(), withCTXEntry.WithFields(logrus.Fields{"test": "test"}))

	bufferWithoutCTX := &bytes.Buffer{}
	logrus.SetOutput(bufferWithoutCTX)

	type args struct {
		ctx  context.Context
		args []interface{}
	}
	tests := []struct {
		name   string
		args   args
		buffer *bytes.Buffer
	}{
		{
			name: "with ctx",
			args: args{
				ctx:  ctxWithEntry,
				args: []interface{}{},
			},
			buffer: bufferForCTX,
		},
		{
			name: "without ctx",
			args: args{
				ctx:  context.Background(),
				args: []interface{}{},
			},
			buffer: bufferWithoutCTX,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.ctx, tt.args.args...)
			if len(tt.buffer.String()) == 0 {
				t.Errorf("Error() = %v, want %v", tt.buffer.String(), tt.args.args)
			}
		})
	}
}

func TestWarnf(t *testing.T) {
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
			Warnf(tt.args.ctx, tt.args.format, tt.args.args...)
			if len(tt.buffer.String()) == 0 {
				t.Errorf("Warnf() = %v, want %v", tt.buffer.String(), tt.args.format)
			}
		})
	}
}

func TestWarn(t *testing.T) {
	bufferForCTX := &bytes.Buffer{}
	withCTXLogger := logrus.New()
	withCTXLogger.SetOutput(bufferForCTX)
	withCTXEntry := logrus.NewEntry(withCTXLogger)

	ctxWithEntry := AddEntryToCTX(context.Background(), withCTXEntry.WithFields(logrus.Fields{"test": "test"}))

	bufferWithoutCTX := &bytes.Buffer{}
	logrus.SetOutput(bufferWithoutCTX)

	type args struct {
		ctx  context.Context
		args []interface{}
	}
	tests := []struct {
		name   string
		args   args
		buffer *bytes.Buffer
	}{
		{
			name: "with ctx",
			args: args{
				ctx:  ctxWithEntry,
				args: []interface{}{},
			},
			buffer: bufferForCTX,
		},
		{
			name: "without ctx",
			args: args{
				ctx:  context.Background(),
				args: []interface{}{},
			},
			buffer: bufferWithoutCTX,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warn(tt.args.ctx, tt.args.args...)
			if len(tt.buffer.String()) == 0 {
				t.Errorf("Warn() = %v, want %v", tt.buffer.String(), tt.args.args)
			}
		})
	}
}

func TestTracef(t *testing.T) {
	bufferForCTX := &bytes.Buffer{}
	withCTXLogger := logrus.New()
	withCTXLogger.SetOutput(bufferForCTX)
	withCTXEntry := logrus.NewEntry(withCTXLogger)
	withCTXLogger.SetLevel(logrus.TraceLevel)

	ctxWithEntry := AddEntryToCTX(context.Background(), withCTXEntry.WithFields(logrus.Fields{"test": "test"}))

	bufferWithoutCTX := &bytes.Buffer{}
	logrus.SetOutput(bufferWithoutCTX)
	logrus.SetLevel(logrus.TraceLevel)

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
			Tracef(tt.args.ctx, tt.args.format, tt.args.args...)
			if len(tt.buffer.String()) == 0 {
				t.Errorf("Tracef() = %v, want %v", tt.buffer.String(), tt.args.format)
			}
		})
	}
}

func TestTrace(t *testing.T) {
	bufferForCTX := &bytes.Buffer{}
	withCTXLogger := logrus.New()
	withCTXLogger.SetOutput(bufferForCTX)
	withCTXEntry := logrus.NewEntry(withCTXLogger)
	withCTXLogger.SetLevel(logrus.TraceLevel)

	ctxWithEntry := AddEntryToCTX(context.Background(), withCTXEntry.WithFields(logrus.Fields{"test": "test"}))

	bufferWithoutCTX := &bytes.Buffer{}
	logrus.SetOutput(bufferWithoutCTX)
	logrus.SetLevel(logrus.TraceLevel)

	type args struct {
		ctx  context.Context
		args []interface{}
	}
	tests := []struct {
		name   string
		args   args
		buffer *bytes.Buffer
	}{
		{
			name: "with ctx",
			args: args{
				ctx:  ctxWithEntry,
				args: []interface{}{},
			},
			buffer: bufferForCTX,
		},
		{
			name: "without ctx",
			args: args{
				ctx:  context.Background(),
				args: []interface{}{},
			},
			buffer: bufferWithoutCTX,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Trace(tt.args.ctx, tt.args.args...)
			if len(tt.buffer.String()) == 0 {
				t.Errorf("Trace() = %v, want %v", tt.buffer.String(), tt.args.args)
			}
		})
	}
}

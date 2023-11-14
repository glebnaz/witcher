package log

import (
	"bytes"
	"context"
	"log/slog"
	"strings"
	"testing"
)

func TestCTXAttrHandler(t *testing.T) {
	type kv struct {
		key   string
		value string
	}

	type args struct {
		kv []kv
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "simple test with one attr in ctx",
			args: args{
				kv: []kv{
					{
						key:   "req-id",
						value: "123",
					},
				},
			},
		},
		{
			name: "simple test with two attr in ctx",
			args: args{
				kv: []kv{
					{
						key:   "req-id",
						value: "123",
					},
					{
						key:   "name",
						value: "John Doe",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			for _, kv := range tt.args.kv {
				ctx = AddSLOGAttrToCTX(ctx, []slog.Attr{slog.String(kv.key, kv.value)})
			}

			buf := &bytes.Buffer{}
			defer buf.Truncate(len(buf.Bytes()))

			handler := NewCTXAttrHandler("test", slog.NewJSONHandler(buf, nil))
			logger := slog.New(handler)
			logger.InfoContext(ctx, "test")
			str := buf.String()

			for _, kv := range tt.args.kv {
				if !strings.Contains(str, kv.key) {
					t.Errorf("JSONCTXHandlerBase() = %v, want %v", str, kv.key)
				}
				if !strings.Contains(str, kv.value) {
					t.Errorf("JSONCTXHandlerBase() = %v, want %v", str, kv.value)
				}
			}
		})
	}
}

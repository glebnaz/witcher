package config

import (
	"context"
	"os"
	"testing"
)

// nolint:paralleltest
func Test_setEnvFromMap(t *testing.T) {
	type args struct {
		ctx         context.Context
		env         map[string]string
		returnError bool
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		wantValue map[string]string
	}{
		{
			name: "simple test",
			args: args{
				ctx: context.Background(),
				env: map[string]string{
					"witcher": "value",
				},
			},
			wantErr: false,
			wantValue: map[string]string{
				"WITCHER": "value",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := setEnvFromMap(tt.args.ctx, tt.args.env, tt.args.returnError); (err != nil) != tt.wantErr {
				t.Errorf("setEnvFromMap() error = %v, wantErr %v", err, tt.wantErr)
			}
			for k, v := range tt.wantValue {
				actual := os.Getenv(k)
				if actual != v {
					t.Errorf("setEnvFromMap() = %v, want %v", actual, v)
				}
			}
		})
	}
}

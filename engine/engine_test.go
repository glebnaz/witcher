package engine

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestDebugServer tests the DebugServer
//
// will check the all k8s probes
func TestDebugServer(t *testing.T) {
	t.Log("Testing DEBUG server")

	type fields struct {
		e        *Server
		waitTime *time.Duration
	}

	type ProbeResult struct {
		StatusCode  int
		Body        string
		CallerError assert.ErrorAssertionFunc
	}

	type ReadyProbeResult struct {
		StatusCode int
		Body       map[string]bool
		CallError  assert.ErrorAssertionFunc
	}

	type want struct {
		Ready        ReadyProbeResult
		Live         ProbeResult
		StartUp      ProbeResult
		RunErrorFunc assert.ErrorAssertionFunc
	}

	tests := []struct {
		name   string
		fields func() fields
		want   want
	}{
		{
			name: "Simple Positive Test(With out checker)",
			fields: func() fields {
				e := NewServer()

				return fields{
					e: e,
				}
			},
			want: want{
				Ready: ReadyProbeResult{
					StatusCode: http.StatusOK,
					Body: map[string]bool{
						"run group": true,
					},
					CallError: assert.NoError,
				},
				Live: ProbeResult{
					StatusCode:  http.StatusOK,
					Body:        "OK",
					CallerError: assert.NoError,
				},
				StartUp: ProbeResult{
					StatusCode:  http.StatusOK,
					Body:        "OK",
					CallerError: assert.NoError,
				},
				RunErrorFunc: assert.NoError,
			},
		},
		{
			name: "Simple Positive Test(With checker)",
			fields: func() fields {
				e := NewServer()

				checker := NewDefaultChecker("checker true", func(ctx context.Context) error {
					return nil
				})

				e.AddCheckers([]Checker{checker})

				time := 1 * time.Second

				return fields{
					e:        e,
					waitTime: &time,
				}
			},
			want: want{
				Ready: ReadyProbeResult{
					StatusCode: http.StatusOK,
					Body: map[string]bool{
						"run group":    true,
						"checker true": true,
					},
					CallError: assert.NoError,
				},
				Live: ProbeResult{
					StatusCode:  http.StatusOK,
					Body:        "OK",
					CallerError: assert.NoError,
				},
				StartUp: ProbeResult{
					StatusCode:  http.StatusOK,
					Body:        "OK",
					CallerError: assert.NoError,
				},
				RunErrorFunc: assert.NoError,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.fields()
			e := f.e
			go func() {
				err := e.Run()
				tt.want.RunErrorFunc(t, err)
			}()

			defer e.Shutdown()

			// wait for the server to start
			if f.waitTime != nil {
				time.Sleep(*f.waitTime)
			}

			// check the ready probe
			resp, err := http.Get("http://localhost:8084/read")
			tt.want.Ready.CallError(t, err)
			if tt.want.Ready.Body != nil {
				bodyByte, err := io.ReadAll(resp.Body)
				assert.NoError(t, err)
				readyData := map[string]bool{}
				err = json.Unmarshal(bodyByte, &readyData)
				assert.NoError(t, err)
				assert.Equal(t, tt.want.Ready.Body, readyData)
				assert.Equal(t, tt.want.Ready.StatusCode, resp.StatusCode)
			}

			// check the live probe
			resp, err = http.Get("http://localhost:8084/live")
			tt.want.Live.CallerError(t, err)
			if tt.want.Live.Body != "" {
				bodyByte, err := io.ReadAll(resp.Body)
				assert.NoError(t, err)
				assert.Equal(t, tt.want.Live.Body, string(bodyByte))
				assert.Equal(t, tt.want.Live.StatusCode, resp.StatusCode)
			}

			// check the startup probe
			resp, err = http.Get("http://localhost:8084/startup")
			tt.want.StartUp.CallerError(t, err)
			if tt.want.StartUp.Body != "" {
				bodyByte, err := io.ReadAll(resp.Body)
				assert.NoError(t, err)
				assert.Equal(t, tt.want.Live.Body, string(bodyByte))
				assert.Equal(t, tt.want.Live.StatusCode, resp.StatusCode)
			}

		})
	}
}

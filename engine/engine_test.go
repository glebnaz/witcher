package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestDebugServer tests the DebugServer
//
// will check the all k8s probes
// nolint:paralleltest
func TestDebugServer(t *testing.T) {
	t.Log("Testing DEBUG server")

	type fields struct {
		e        *Server
		waitTime *time.Duration
		ctx      context.Context
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
				e := NewServer(WithDebugPort(":1111"))

				ctx := context.Background()

				return fields{
					e:   e,
					ctx: ctx,
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
				e := NewServer(WithDebugPort(":1111"))

				checker := NewDefaultChecker("checker true", func(ctx context.Context) error {
					return nil
				})

				e.AddCheckers([]Checker{checker})

				time := 1 * time.Second

				ctx := context.Background()

				return fields{
					e:        e,
					waitTime: &time,
					ctx:      ctx,
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
		{
			name: "Simple Negative Test(With checker)",
			fields: func() fields {
				e := NewServer(WithDebugPort(":1111"))

				checker := NewDefaultChecker("checker false", func(ctx context.Context) error {
					return assert.AnError
				})

				e.AddCheckers([]Checker{checker})

				time := 1 * time.Second

				ctx := context.Background()

				return fields{
					e:        e,
					waitTime: &time,
					ctx:      ctx,
				}
			},
			want: want{
				Ready: ReadyProbeResult{
					StatusCode: http.StatusInternalServerError,
					Body: map[string]bool{
						"run group":     true,
						"checker false": false,
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

	// nolint:paralleltest
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.fields()
			e := f.e
			go func() {
				err := e.Run()
				tt.want.RunErrorFunc(t, err)
			}()

			t.Cleanup(func() {
				errShutdown := e.Shutdown()
				assert.NoError(t, errShutdown)
			})

			// wait for the server to start
			if f.waitTime != nil {
				time.Sleep(*f.waitTime)
			}

			// check the ready probe
			resp, err := makeRequest(f.ctx, "http://localhost:1111/read")
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
			resp, err = makeRequest(f.ctx, "http://localhost:1111/live")
			tt.want.Live.CallerError(t, err)
			if tt.want.Live.Body != "" {
				bodyByte, err := io.ReadAll(resp.Body)
				assert.NoError(t, err)
				assert.Equal(t, tt.want.Live.Body, string(bodyByte))
				assert.Equal(t, tt.want.Live.StatusCode, resp.StatusCode)
			}

			// check the startup probe
			resp, err = makeRequest(f.ctx, "http://localhost:1111/startup")
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

func makeRequest(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

const (
	liveHandlerPostfix = "/live"
	host               = "http://localhost"
)

func TestWithDebugPort(t *testing.T) {
	t.Parallel()
	//get random int from 1000 to 9999
	port := 1000 + time.Now().Nanosecond()%9000
	portString := fmt.Sprintf(":%d", port)
	s := NewServer(WithDebugPort(portString))
	assert.Equal(t, portString, s.DebugServer.PORT)

	go func() {
		err := s.Run()
		assert.NoError(t, err)
	}()

	time.Sleep(5 * time.Millisecond)

	resp, err := makeRequest(context.Background(), host+portString+liveHandlerPostfix)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDebugPortFromEnv(t *testing.T) {
	t.Parallel()
	time.Sleep(5 * time.Millisecond)
	port := 1000 + time.Now().Nanosecond()%9000
	portString := fmt.Sprintf(":%d", port)

	//set env
	err := os.Setenv("DEBUG_PORT", portString)
	assert.NoError(t, err)
	s := NewServer()
	go func() {
		err := s.Run()
		assert.NoError(t, err)
	}()

	time.Sleep(5 * time.Millisecond)

	resp, err := makeRequest(context.Background(), host+portString+liveHandlerPostfix)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestServer_AddCloser(t *testing.T) {
	t.Parallel()
	s := NewServer()
	closer := NewDefaultCloser("closer", func(ctx context.Context) error {
		return nil
	})
	s.AddCloser(closer)
	assert.Equal(t, 1, len(s.closerGroup))
	assert.Equal(t, "closer", s.closerGroup[0].GetName())
}

func TestServer_AddChecker(t *testing.T) {
	t.Parallel()
	s := NewServer()
	checker := NewDefaultChecker("checker", func(ctx context.Context) error {
		return nil
	})
	s.AddChecker(checker)
	assert.Equal(t, 1, len(s.checkersGroup))
	assert.Equal(t, "checker", s.checkersGroup[0].GetName())
}

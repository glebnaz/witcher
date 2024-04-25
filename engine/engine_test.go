package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestDebugServer tests the DebugServer
//
// will check the all k8s probes
func TestDebugServer(t *testing.T) {
	t.Parallel()
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
		fields func(p string) fields
		Port   string
		want   want
	}{
		{
			name: "Simple Positive Test(With out checker)",
			fields: func(p string) fields {
				e := NewServer(WithDebugPort(p))

				ctx := context.Background()

				return fields{
					e:   e,
					ctx: ctx,
				}
			},
			Port: getUniquePort(),
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
			Port: getUniquePort(),
			fields: func(p string) fields {
				e := NewServer(WithDebugPort(p))

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
			Port: getUniquePort(),
			fields: func(p string) fields {
				e := NewServer(WithDebugPort(p))

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

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := tt.fields(tt.Port)
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
			resp, err := makeRequest(f.ctx, host+tt.Port+"/read")
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
			resp, err = makeRequest(f.ctx, host+tt.Port+"/live")
			tt.want.Live.CallerError(t, err)
			if tt.want.Live.Body != "" {
				bodyByte, err := io.ReadAll(resp.Body)
				assert.NoError(t, err)
				assert.Equal(t, tt.want.Live.Body, string(bodyByte))
				assert.Equal(t, tt.want.Live.StatusCode, resp.StatusCode)
			}

			// check the startup probe
			resp, err = makeRequest(f.ctx, host+tt.Port+"/startup")
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

type portContainer struct {
	last int
	m    sync.Mutex
}

var pContainer = portContainer{
	last: 3500, // use this values because gh actions restricts the port range
}

func getUniquePort() string {
	pContainer.m.Lock()
	last := pContainer.last
	newPort := last + 1
	pContainer.last = newPort
	pContainer.m.Unlock()
	return fmt.Sprintf(":%d", newPort)
}

func TestWithDebugPort(t *testing.T) {
	t.Parallel()
	//get random int from 1000 to 9999
	portString := getUniquePort()
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
	portString := getUniquePort()

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

func TestServer_RunGroup(t *testing.T) {
	t.Parallel()
	port := getUniquePort()
	s := NewServer(WithDebugPort(port))

	data := make(chan int)

	s.AddActor(func() error {
		for i := 0; i < 10; i++ {
			data <- i
		}
		err := s.Shutdown()
		assert.NoError(t, err)
		return nil
	}, func(error) {
		assert.NoError(t, nil)
	})

	go func() {
		err := s.Run()
		assert.NoError(t, err)
	}()

	value := make([]int, 0, 10)
	for {
		i := <-data
		value = append(value, i)

		if len(value) == 10 {
			break
		}
	}

	assert.Equal(t, 10, len(value))
}

func TestPProfHandlers(t *testing.T) {
	t.Parallel()
	mux := http.NewServeMux()
	wrapPProf(mux)

	tests := []struct {
		name string
		path string
	}{
		{"Index", "/debug/pprof/"},
		{"Heap", "/debug/pprof/heap"},
		{"Goroutine", "/debug/pprof/goroutine"},
		{"Block", "/debug/pprof/block"},
		{"ThreadCreate", "/debug/pprof/threadcreate"},
		{"Cmdline", "/debug/pprof/cmdline"},
		{"Profile", "/debug/pprof/profile"},
		{"Symbol", "/debug/pprof/symbol"},
		{"Trace", "/debug/pprof/trace"},
		{"Mutex", "/debug/pprof/mutex"},
		{"Allocs", "/debug/pprof/allocs"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			//nolint:noctx
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusOK, rr.Code)
		})
	}
}

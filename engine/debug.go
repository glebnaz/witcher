package engine

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/glebnaz/witcher/metrics"
	"github.com/rs/zerolog/log"
)

type DebugServer struct {
	PORT string `json:"port" envconfig:"PORT" default:":8084"`

	mux    *http.ServeMux
	server *http.Server

	m sync.RWMutex

	ready         bool
	checkersGroup []Checker
}

var readHeaderTimeout = 2 * time.Minute

// NewDebugServer create new debug server
//
// port is debug port
// use value with `:` for example: :8084
func NewDebugServer(port string) *DebugServer {
	debug := &DebugServer{
		PORT:          port,
		ready:         false,
		checkersGroup: make([]Checker, 0, 10),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/read", debug.ReadinessProbeHandler)
	mux.HandleFunc("/live", debug.LiveProbeHandler)
	mux.HandleFunc("/startup", debug.StartupProbeHandler)
	mux.Handle("/metrics", metrics.Handler())
	wrapPProf(mux)

	debug.mux = mux

	debug.server = &http.Server{
		Addr:              debug.PORT,
		Handler:           mux,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	return debug
}

// SetReady set server ready
//
// warning use this method only when you sure that server is ready
func (d *DebugServer) SetReady(ready bool) {
	d.m.Lock()
	defer d.m.Unlock()
	d.ready = ready
	if ready {
		log.Info().Msg("Server set ready")
	} else {
		log.Info().Msg("Server is not ready")
	}
}

func (d *DebugServer) AddChecker(checker Checker) {
	log.Debug().Msgf("Adding checker %s", checker.GetName())
	d.m.Lock()
	defer d.m.Unlock()
	d.checkersGroup = append(d.checkersGroup, checker)
}

// AddCheckers for check your server is live
func (d *DebugServer) AddCheckers(checkers []Checker) {
	for _, checker := range checkers {
		d.AddChecker(checker)
	}
}

// ReadinessProbeHandler is probe checker for k8s readiness probe
//
// It will check all checkersGroup and return 200 if all checkersGroup is ok
// or 500 if one of them is not ok
//
// response will be json with all checkersGroup and their status
func (d *DebugServer) ReadinessProbeHandler(w http.ResponseWriter, req *http.Request) {
	log.Info().Msgf("readiness probe at %s", time.Now())
	d.m.RLock()
	defer d.m.RUnlock()

	info := make(map[string]bool)

	live := true

	for i := range d.checkersGroup {
		if err := d.checkersGroup[i].Check(req.Context()); err != nil {
			log.Error().Msgf("Server is not ready to receive traffic: %s", err)
			live = false
			info[d.checkersGroup[i].GetName()] = false
		} else {
			info[d.checkersGroup[i].GetName()] = true
		}
	}

	jsonInfo, err := json.Marshal(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !live {
		log.Error().Msgf("Server is not ready to receive traffic")
		w.WriteHeader(http.StatusInternalServerError)
		//nolint:errcheck
		w.Write(jsonInfo)
		return
	}
	w.WriteHeader(http.StatusOK)
	//nolint:errcheck
	w.Write(jsonInfo)
}

// LiveProbeHandler is probe checker for k8s liveness probe
func (d *DebugServer) LiveProbeHandler(w http.ResponseWriter, _ *http.Request) {
	log.Info().Msgf("live probe at %s", time.Now())
	w.WriteHeader(http.StatusOK)
	//nolint:errcheck
	w.Write([]byte("OK"))
}

// StartupProbeHandler is probe checker for k8s startup probe
// ready or not ready
// used only for start
func (d *DebugServer) StartupProbeHandler(w http.ResponseWriter, _ *http.Request) {
	log.Info().Msgf("StartupProbe check at %s", time.Now())
	if d.ready {
		w.WriteHeader(http.StatusOK)
		// nolint:errcheck
		w.Write([]byte("OK"))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		// nolint:errcheck
		w.Write([]byte("Not ready"))
	}
}

func (d *DebugServer) RunDebug() error {
	d.server.Addr = d.PORT
	log.Info().Msgf("Run debug server at %s", time.Now())
	return d.server.ListenAndServe()
}

func (d *DebugServer) ShutdownDebug(ctx context.Context) error {
	log.Info().Msgf("Start shutdown debug server at %s", time.Now())
	errShutDown := d.server.Shutdown(ctx)
	if errShutDown != nil {
		return errShutDown
	}
	log.Debug().Msg("Shutdown debug server success")
	return nil
}

package engine

import (
	"context"
	"github.com/glebnaz/witcher/swagger"
	"net/http"
	"sync"
	"time"

	"github.com/glebnaz/witcher/metrics"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type DebugServer struct {
	PORT string `json:"port" envconfig:"PORT" default:":8084"`

	engine *echo.Echo
	m      sync.Mutex

	ready    bool
	checkers []Checker
}

func NewDebugServer(port string) *DebugServer {
	e := echo.New()
	e.Debug = false
	e.HideBanner = true
	e.HidePort = true

	debug := &DebugServer{
		PORT:     port,
		ready:    false,
		checkers: make([]Checker, 0, 10),
	}

	e.GET("/ready", debug.Ready)
	e.GET("/live", debug.Live)
	e.GET("/metrics", echo.WrapHandler(metrics.Handler()))

	//file, err := os.Open("/Users/glebnaz/Documents/#main/workspace/go-platform/swagger/swagger.json")
	//fmt.Println(err)
	//data, err := io.ReadAll(file)
	//fmt.Println(err)
	err := swagger.AddSwagger(e, "/Users/glebnaz/Documents/#main/workspace/go-platform/swagger.json")
	if err != nil {
		log.Errorf("Error adding swagger: %s", err)
	}
	wrapPProf(e)
	debug.engine = e

	return debug
}

//SetReady set server ready
//
// warning use this method only when you sure that server is ready
func (d *DebugServer) SetReady(ready bool) {
	d.m.Lock()
	defer d.m.Unlock()
	d.ready = ready
	if ready {
		log.Infof("Server is ready")
	} else {
		log.Infof("Server is not ready")
	}
}

func (d *DebugServer) AddChecker(checker Checker) {
	log.Debugf("Adding checker %s", checker.Name())
	d.m.Lock()
	defer d.m.Unlock()
	d.checkers = append(d.checkers, checker)
}

//AddCheckers for check your server is live
func (d *DebugServer) AddCheckers(checkers []Checker) {
	for _, checker := range checkers {
		d.AddChecker(checker)
	}
}

//Live is probe checker
func (d *DebugServer) Live(c echo.Context) error {
	log.Infof("Live check at %s", time.Now())
	d.m.Lock()
	defer d.m.Unlock()

	info := make(map[string]bool)

	live := true

	for i := range d.checkers {
		if err := d.checkers[i].Check(); err != nil {
			log.Errorf("Server is not live: %s", err)
			live = false
			info[d.checkers[i].Name()] = false
		} else {
			info[d.checkers[i].Name()] = true
		}
	}

	if !live {
		log.Errorf("Server is not live")
		return c.JSON(http.StatusInternalServerError, info)
	}
	return c.JSON(http.StatusOK, info)
}

//Ready is probe checker
func (d *DebugServer) Ready(c echo.Context) error {
	log.Infof("Ready check at %s", time.Now())
	if d.ready {
		return c.String(http.StatusOK, "OK")
	}
	return c.String(http.StatusInternalServerError, "Not ready")
}

func (d *DebugServer) RunDebug() error {
	log.Infof("Run debug server at %s", time.Now())
	return d.engine.Start(d.PORT)
}

func (d *DebugServer) ShutdownDebug(ctx context.Context) error {
	log.Infof("Start shutdown debug server at %s", time.Now())
	errShutDown := d.engine.Shutdown(ctx)
	if errShutDown != nil {
		return errShutDown
	}
	log.Debugf("Shutdown debug server success")
	return nil
}

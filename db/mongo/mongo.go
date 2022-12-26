package mongo

import (
	"context"
	"sync"
	"time"

	"github.com/glebnaz/witcher/engine"
	log "github.com/sirupsen/logrus"

	driver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	*driver.Client

	lock sync.RWMutex

	name string
}

// HealthChecker use for registration mongo health check in wither engine
// timeout is duration for ping (use context.WithTimeout inside)
func (m *Mongo) HealthChecker(timeout time.Duration) engine.Checker {
	return engine.NewDefaultChecker(m.GetName(), func() error {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		err := m.Ping(ctx, nil)
		defer cancel()
		if err != nil {
			return err
		}
		return nil
	})
}

// Closer use for registration mongo closer in wither engine
func (m *Mongo) Closer() engine.Closer {
	return engine.NewDefaultCloser(m.GetName(), func(ctx context.Context, group *sync.WaitGroup) error {
		defer group.Done()
		err := m.Disconnect(ctx)
		if err != nil {
			log.Debugf("Error disconnect mongo: %s", err)
			return err
		}
		return nil
	})
}

func (m *Mongo) GetName() string {
	return m.name
}

// ChangeInstanceName use this method if you want change closer and checker name
func (m *Mongo) ChangeInstanceName(name string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.name = name
}

// NewMongo returns new instance of Mongo
// as driver wither use go.mongodb.org/mongo-driver/mongo
//
// ctx - use ctx from engine function GetCTX()
// opts is ClientOptions from go.mongodb.org/mongo-driver/mongo/options
func NewMongo(ctx context.Context, uri string, opts ...*options.ClientOptions) (*Mongo, error) {
	opts = append(opts, options.Client().ApplyURI(uri))
	conn, err := driver.Connect(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return &Mongo{
		Client: conn,
		name:   "mongo",
	}, nil
}

package mongo

import (
	"context"
	"testing"
	"time"

	"github.com/glebnaz/witcher/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type MongoTestSuite struct {
	suite.Suite
	container testcontainers.Container

	endpoint string
}

const (
	uriProtocol = "mongodb://"
)

func (m *MongoTestSuite) SetupTest() {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForExposedPort(),
	}

	l := log.NewTestLogger(log.WithTestLoggerName("mongo_test_suite"))

	mongoContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
		Logger:           l,
	})

	if err != nil {
		m.FailNow(err.Error())
	}

	m.container = mongoContainer
	endpoint, err := m.container.Endpoint(ctx, "")
	if err != nil {
		m.FailNow(err.Error())
	}
	m.endpoint = uriProtocol + endpoint
}

func (m *MongoTestSuite) TearDownTest() {
	err := m.container.Terminate(context.Background())
	if err != nil {
		m.FailNow(err.Error())
	}
}

func (m *MongoTestSuite) TestCloser() {
	t := m.Suite.T()

	type args struct {
		ctx    context.Context
		cancel context.CancelFunc
		uri    string
	}

	type want struct {
		errFunc assert.ErrorAssertionFunc
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success close",
			args: args{
				ctx: context.Background(),
				uri: m.endpoint,
			},
			want: want{
				errFunc: assert.NoError,
			},
		},
	}

	for _, tt := range tests {
		m.Run(tt.name, func() {
			instance, err := NewMongo(tt.args.ctx, tt.args.uri)
			assert.NoError(t, err)
			closer := instance.Closer()

			if tt.args.cancel != nil {
				tt.args.cancel()
			}

			err = closer.Close(tt.args.ctx)
			tt.want.errFunc(t, err)
		})
	}
}

func (m *MongoTestSuite) TestHealthChecker() {
	t := m.Suite.T()

	type args struct {
		ctx    context.Context
		cancel context.CancelFunc
		uri    string
	}

	type want struct {
		errFunc assert.ErrorAssertionFunc
	}

	ctxForCancl, cancel := context.WithCancel(context.Background())

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success check",
			args: args{
				ctx:    context.Background(),
				uri:    m.endpoint,
				cancel: nil,
			},
			want: want{
				errFunc: assert.NoError,
			},
		},
		{
			name: "cancel context",
			args: args{
				ctx:    ctxForCancl,
				cancel: cancel,
				uri:    "mongodb://" + m.endpoint,
			},
			want: want{
				errFunc: assert.Error,
			},
		},
	}

	for _, tt := range tests {
		m.Run(tt.name, func() {
			tt := tt
			instance, err := NewMongo(tt.args.ctx, tt.args.uri)
			assert.NoError(t, err)
			checker := instance.HealthChecker(5 * time.Second)

			if tt.args.cancel != nil {
				tt.args.cancel()
			}

			err = checker.Check(tt.args.ctx)
			tt.want.errFunc(t, err)
		})
	}
}

func TestMongoTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(MongoTestSuite))
}

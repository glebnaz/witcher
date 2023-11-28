package engine

import (
	"context"
	"net"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type grpcServerRunner struct {
	port                    string
	server                  *grpc.Server
	gracefulShutdownTimeout time.Duration
}

func newGRPCServerRunner(grpcServer *grpc.Server, port string, gracefulTimeout time.Duration) *grpcServerRunner {
	return &grpcServerRunner{
		port:                    port,
		server:                  grpcServer,
		gracefulShutdownTimeout: gracefulTimeout,
	}
}

func (b *grpcServerRunner) actor() (func() error, func(error)) {
	return func() error {
			listener, err := net.Listen("tcp", b.port)
			if err != nil {
				return errors.Wrap(err, "grpc listener init failure")
			}
			log.Info().Msgf("grpc server: started on %s port", b.port)

			err = errors.Wrap(b.server.Serve(listener), "grpc server")
			if err != nil {
				log.Error().Msgf("grpc serve error: %s", err)
			}
			return err
		}, func(err error) {
			doneCh := make(chan struct{})
			go func() {
				b.server.GracefulStop()
				close(doneCh)
			}()

			select {
			case <-time.After(b.gracefulShutdownTimeout):
				log.Error().Err(errors.Wrap(context.DeadlineExceeded, "grpc server graceful stop timed out")).Msg("stop server by timeout")
				b.server.Stop()
				log.Info().Msg("grpc server stopped (force)")
			case <-doneCh:
				log.Info().Msg("grpc server: gracefully stopped")
			}
		}
}

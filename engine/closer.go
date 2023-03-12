package engine

import (
	"context"
	"sync"
)

type Closer interface {
	GetName() string

	Close(ctx context.Context, wg *sync.WaitGroup) error
}

// DefaultCloser is a default implementation of Closer
// if you want simple checker without implementation Closer interface
// tou can call this function and receive Closer interface
type DefaultCloser struct {
	Name      string                                              `json:"name"`
	CloseFunc func(ctx context.Context, wg *sync.WaitGroup) error `json:"close"`
}

func (d *DefaultCloser) GetName() string {
	return d.Name
}

func (d *DefaultCloser) Close(ctx context.Context, wg *sync.WaitGroup) error {
	return d.CloseFunc(ctx, wg)
}

func NewDefaultCloser(name string, close func(ctx context.Context, group *sync.WaitGroup) error) Closer {
	return &DefaultCloser{
		Name:      name,
		CloseFunc: close,
	}
}

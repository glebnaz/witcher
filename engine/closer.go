package engine

import (
	"context"
)

type Closer interface {
	GetName() string

	Close(ctx context.Context) error
}

// DefaultCloser is a default implementation of Closer
// if you want simple checker without implementation Closer interface
// tou can call this function and receive Closer interface
type DefaultCloser struct {
	Name      string                          `json:"name"`
	CloseFunc func(ctx context.Context) error `json:"close"`
}

func (d *DefaultCloser) GetName() string {
	return d.Name
}

func (d *DefaultCloser) Close(ctx context.Context) error {
	return d.CloseFunc(ctx)
}

func NewDefaultCloser(name string, close func(ctx context.Context) error) Closer {
	return &DefaultCloser{
		Name:      name,
		CloseFunc: close,
	}
}

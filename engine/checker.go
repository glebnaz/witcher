package engine

import "context"

type Checker interface {
	Check(ctx context.Context) error
	GetName() string
}

// DefaultChecker is a default implementation of Checker
// if you want simple checker without implementation Checker interface
// tou can call this function and receive Checker interface
type DefaultChecker struct {
	CheckFunc func(ctx context.Context) error `json:"check"`
	NameCheck string                          `json:"name"`
}

func (c *DefaultChecker) Check(ctx context.Context) error {
	return c.CheckFunc(ctx)
}

func (c *DefaultChecker) GetName() string {
	return c.NameCheck
}

func NewDefaultChecker(name string, checkFunc func(ctx context.Context) error) *DefaultChecker {
	return &DefaultChecker{
		CheckFunc: checkFunc,
		NameCheck: name,
	}
}

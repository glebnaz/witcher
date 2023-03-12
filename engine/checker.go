package engine

type Checker interface {
	Check() error
	Name() string
}

// DefaultChecker is a default implementation of Checker
// if you want simple checker without implementation Checker interface
// tou can call this function and receive Checker interface
type DefaultChecker struct {
	CheckFunc func() error `json:"check"`
	NameCheck string       `json:"name"`
}

func (c *DefaultChecker) Check() error {
	return c.CheckFunc()
}

func (c *DefaultChecker) Name() string {
	return c.NameCheck
}

func NewDefaultChecker(name string, checkFunc func() error) *DefaultChecker {
	return &DefaultChecker{
		CheckFunc: checkFunc,
		NameCheck: name,
	}
}

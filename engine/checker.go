package engine

type Checker interface {
	Check() error
	Name() string
}

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

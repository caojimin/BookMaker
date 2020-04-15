package book

type Checker struct {
	checkers []func() error
}

func NewChecker(checkers ...func() error) *Checker {
	return &Checker{
		checkers,
	}
}

func (c *Checker) Check() error {
	var err error
	for _, f := range c.checkers {
		if err = f(); err != nil {
			return err
		}
	}
	return nil
}

func GenFileChecker() error {
	_, err := getGenPath()
	return err
}

var DefaultPreChecker = NewChecker(GenFileChecker)

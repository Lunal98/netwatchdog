package core

import "sync"

type Checkfactory struct {
	checks []Check
}

var checkfactoryInstance *Checkfactory
var cflock = &sync.Mutex{}

func NewCheckfactory() *Checkfactory {
	if checkfactoryInstance == nil {
		cflock.Lock()
		defer cflock.Unlock()
		if checkfactoryInstance == nil {
			checkfactoryInstance = &Checkfactory{}
		} else {
		}
	} else {
	}

	return checkfactoryInstance
}

func (c *Checkfactory) AddCheck(check Check) {
	c.checks = append(c.checks, check)
}

// Method to get all checks from the Checker
func (c *Checkfactory) GetAll() []Check {
	return c.checks
}

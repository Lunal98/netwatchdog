package remediationhandler

import (
	"sync"

	"github.com/Lunal98/netwatchdog/internal/remediationhelper"
)

type remediationhandler struct {
	remediationhelper remediationhelper.Helper
	once              sync.Once
}

func (handler *remediationhandler) init() {
	handler.once.Do(func() {
		handler.remediationhelper = remediationhelper.GetDebugHelper()
	})
}
func (handler *remediationhandler) handle() {

}

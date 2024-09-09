package remediationhandler

import (
	"github.com/Lunal98/netwatchdog/internal/checker"
	"github.com/Lunal98/netwatchdog/internal/remediationhelper"
	"github.com/google/uuid"
)

/*
type remediationhandler struct {
	remediationhelper remediationhelper.Helper
	once              sync.Once
}

func (handler *remediationhandler) init() {
	handler.once.Do(func() {
		handler.remediationhelper = remediationhelper.GetDebugHelper()
	})
}
func (handler *remediationhandler) handle(jobID uuid.UUID, jobName string, err error) {

}
*/

func Handle(jobID uuid.UUID, jobName string, err error, checkers map[uuid.UUID]checker.Checker) {
	remhelp := remediationhelper.GetDebugHelper()

}

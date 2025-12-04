package v1

import (
	"github.com/airats/bug_and_fix/internal/service"
	bugandfix "github.com/airats/bug_and_fix/internal/service/bug_and_fix"
)

// StrictServerInterface
type BugAndFixV1 struct {
	bugAndFixService service.BugAndFixService
}

func NewBugAndFixV1() *BugAndFixV1 {
	return &BugAndFixV1{
		bugAndFixService: bugandfix.NewService(),
	}
}

package v1

import (
	"context"

	api_v1 "github.com/airats/bug_and_fix/pkg/api/rest"
)

func (s *BugAndFixV1) GetBugById(ctx context.Context, request api_v1.GetBugByIdRequestObject) (api_v1.GetBugByIdResponseObject, error) {
	bug := api_v1.Bug{
		Id:          12,
		Title:       "Some bug",
		Description: "Some desc",
		Solution:    nil,
	}

	return api_v1.GetBugById200JSONResponse(bug), nil
}

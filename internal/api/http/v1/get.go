package v1

import (
	"context"

	api_v1 "github.com/airats/bug_and_fix/pkg/api/rest"
)

func (s *BugAndFixV1) GetBugById(ctx context.Context, request api_v1.GetBugByIdRequestObject) (api_v1.GetBugByIdResponseObject, error) {
	bug, err := s.bugAndFixService.GetBugById(request.Id)
	if err != nil {
		return api_v1.GetBugById404Response{}, nil
	}

	respBug := api_v1.Bug{
		Id:          bug.ID,
		Title:       bug.Title,
		Description: bug.Description,
		Solution:    nil,
	}

	return api_v1.GetBugById200JSONResponse(respBug), nil
}

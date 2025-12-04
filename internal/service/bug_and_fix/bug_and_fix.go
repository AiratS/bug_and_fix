package bugandfix

import "github.com/airats/bug_and_fix/internal/model"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetBugById(id int) (model.Bug, error) {
	// TODO: call database
	return model.Bug{
		ID:          45,
		Title:       "Hello",
		Description: "Eap",
	}, nil
}

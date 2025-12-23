package bugandfix

import (
	"github.com/airats/bug_and_fix/internal/model"
	repo "github.com/airats/bug_and_fix/internal/repository"
)

type Service struct {
	bugRepository repo.BugRespository
}

func NewService(bugRepository repo.BugRespository) *Service {
	return &Service{
		bugRepository,
	}
}

func (s *Service) GetBugById(id int) (*model.Bug, error) {
	bug, err := s.bugRepository.Get(id)
	if err != nil {
		return nil, err
	}

	return bug, nil
}

package service

import "github.com/airats/bug_and_fix/internal/model"

type BugAndFixService interface {
	GetBugById(id int) (*model.Bug, error)
}

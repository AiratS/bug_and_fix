package repository

import "github.com/airats/bug_and_fix/internal/model"

type BugRespository interface {
	Get(id int) (*model.Bug, error)
}

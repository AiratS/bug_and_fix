package model

import "time"

type Bug struct {
	ID           int        `db:"id"`
	Title        string     `db:"title"`
	Description  string     `db:"description"`
	Project      string     `db:"project"`
	ErrorMessage *string    `db:"error_message"`
	Device       string     `db:"device"`
	Env          int        `db:"env"`
	Solution     *string    `db:"solution"`
	CreatedAt    *time.Time `db:"created_at"`
	SolvedAt     *time.Time `db:"solved_at"`
}

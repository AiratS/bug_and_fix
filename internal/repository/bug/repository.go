package bug

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/airats/bug_and_fix/internal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type repository struct {
	dbPool *pgxpool.Pool
}

func NewBugRepository(dbPool *pgxpool.Pool) *repository {
	return &repository{
		dbPool,
	}
}

func (r *repository) Get(id int) (*model.Bug, error) {
	bugs := psql.Select("*").
		From("bugs").
		Where(sq.Eq{"project": "BugAndFix"})

	sql, args, err := bugs.ToSql()
	fmt.Println(sql, args)
	if err != nil {
		return nil, err
	}

	rows, err := r.dbPool.Query(context.Background(), sql, args...)
	var bug model.Bug
	err = pgxscan.ScanOne(&bug, rows)
	if err != nil {
		return nil, err
	}

	return &bug, nil
}

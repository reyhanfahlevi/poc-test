package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/reyhanfahlevi/poc-test/svc/account/internal/entity"
)

// Repo struct
type Repo struct {
	db *sqlx.DB
}

// New will instantiate repo instance
func New(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) getError(err error) error {
	switch true {
	case strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint"):
		return entity.ErrIsExist
	default:
		return err
	}
}

// Get get users using specified param
func (r *Repo) Get(ctx context.Context, param entity.ParamGetUser) ([]entity.User, error) {
	var (
		users = []entity.User{}
		where = entity.GetDBParam(param).GenerateNamedWhereQueryFormula()
		arg   = entity.GetDBParam(param).GenerateNamedQueryArgs()
	)

	q := fmt.Sprintf(`SELECT * FROM wp_user %s LIMIT :limit OFFSET :offset`, where)

	query, args, _ := sqlx.Named(q, arg)
	query, args, _ = sqlx.In(query, args...)
	err := r.db.SelectContext(ctx, &users, r.db.Rebind(query), args...)
	if err != nil {
		return users, r.getError(err)
	}

	return users, nil
}

func (r *Repo) Save(ctx context.Context, data entity.ParamSaveUser) error {
	var (
		id int64
	)

	q := `INSERT INTO 
			wp_user (
				name, 
				email, 
				pwd_hash,
				reset_token,
				created_time, 
				created_by, 
				updated_time, 
				updated_by
			) VALUES (
				?, ?, ?, '', now(), ?, now(), ?
			) RETURNING id`

	row := r.db.QueryRowContext(
		ctx,
		r.db.Rebind(q),
		data.Name,
		data.Email,
		data.PWDHash,
		data.CreatedBy,
		data.CreatedBy)

	err := row.Scan(&id)
	if err != nil {
		return r.getError(err)
	}

	return nil
}

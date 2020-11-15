package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/reyhanfahlevi/poc-test/svc/shop/internal/entity"
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

// Get get shops using specified param
func (r *Repo) Get(ctx context.Context, param entity.ParamGetShop) ([]entity.Shop, error) {
	var (
		shops = []entity.Shop{}
		where = entity.GetDBParam(param).GenerateNamedWhereQueryFormula()
		arg   = entity.GetDBParam(param).GenerateNamedQueryArgs()
	)

	q := fmt.Sprintf(`SELECT * FROM wp_shop %s LIMIT :limit OFFSET :offset`, where)

	query, args, _ := sqlx.Named(q, arg)
	query, args, _ = sqlx.In(query, args...)
	err := r.db.SelectContext(ctx, &shops, r.db.Rebind(query), args...)
	if err != nil {
		return shops, r.getError(err)
	}

	return shops, nil
}

func (r *Repo) Save(ctx context.Context, data entity.ParamSaveShop) error {
	var (
		id int64
	)

	q := `INSERT INTO 
			wp_shop (
				name, 
				user_id, 
				address,
				created_time, 
				created_by, 
				updated_time, 
				updated_by
			) VALUES (
				?, ?, ?, now(), ?, now(), ?
			) RETURNING id`

	row := r.db.QueryRowContext(
		ctx,
		r.db.Rebind(q),
		data.Name,
		data.UserID,
		data.Address,
		data.CreatedBy,
		data.CreatedBy)

	err := row.Scan(&id)
	if err != nil {
		return r.getError(err)
	}

	return nil
}

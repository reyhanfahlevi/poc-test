package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/reyhanfahlevi/poc-test/svc/product/internal/entity"
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

// Get get products using specified param
func (r *Repo) Get(ctx context.Context, param entity.ParamGetProduct) ([]entity.Product, error) {
	var (
		products = []entity.Product{}
		where    = entity.GetDBParam(param).GenerateNamedWhereQueryFormula()
		arg      = entity.GetDBParam(param).GenerateNamedQueryArgs()
	)

	q := fmt.Sprintf(`SELECT * FROM wp_product %s LIMIT :limit OFFSET :offset`, where)

	query, args, _ := sqlx.Named(q, arg)
	query, args, _ = sqlx.In(query, args...)
	err := r.db.SelectContext(ctx, &products, r.db.Rebind(query), args...)
	if err != nil {
		return products, r.getError(err)
	}

	return products, nil
}

func (r *Repo) GetProductImage(ctx context.Context, productID int64) ([]entity.ProductImg, error) {
	var (
		images = []entity.ProductImg{}
		arg    = map[string]interface{}{
			"product_id": productID,
		}
	)

	q := fmt.Sprintf(`SELECT * FROM wp_product_img WHERE product_id = :product_id LIMIT 10`)

	query, args, _ := sqlx.Named(q, arg)
	query, args, _ = sqlx.In(query, args...)
	err := r.db.SelectContext(ctx, &images, r.db.Rebind(query), args...)
	if err != nil {
		return images, r.getError(err)
	}

	return images, nil
}

func (r *Repo) Save(ctx context.Context, data entity.ParamSaveProduct) error {
	var (
		id  int64
		err error
	)

	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	qInsert := `INSERT INTO
			wp_product (
				name, 
				shop_id, 
				description,
				price,
				created_time, 
				created_by, 
				updated_time, 
				updated_by
			) VALUES (
				?, ?, ?, ?, now(), ?, now(), ?
			) RETURNING id`

	row := tx.QueryRowContext(
		ctx,
		r.db.Rebind(qInsert),
		data.Name,
		data.ShopID,
		data.Description,
		data.Price,
		data.CreatedBy,
		data.CreatedBy)

	err = row.Scan(&id)
	if err != nil {
		return r.getError(err)
	}

	qImg := `INSERT INTO wp_product_img (
    			    product_id, 
					image_url, 
    			    created_time, 
    			    created_by, 
    			    updated_time, 
					updated_by
				) VALUES (
					  ?, ?, now(), ?, now(), ?
			  	)`

	for _, img := range data.Images {
		_, err := tx.ExecContext(ctx, tx.Rebind(qImg), id, img.ImageURL, img.CreatedBy, img.CreatedBy)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *Repo) UpdateProduct(ctx context.Context, param entity.ParamSaveProduct) error {
	qUpdate := `UPDATE wp_product SET
					name = ?,
					price = ?,
					description = ?,
                    status = ?,
					updated_by = ?,
					updated_time = ?
				WHERE id = ? and shop_id = ?`

	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	_, err = tx.ExecContext(ctx, tx.Rebind(qUpdate), param.Name, param.Price, param.Description, param.Status, param.UpdatedBy, param.UpdatedTime, param.ID, param.ShopID)
	if err != nil {
		return err
	}

	if len(param.Images) > 0 {
		err = r.updateProductImage(ctx, tx, param.ID, param.Images)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *Repo) UpdateProductImage(ctx context.Context, productID int64, images []entity.ProductImg) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	err = r.updateProductImage(ctx, tx, productID, images)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *Repo) cleanProductImage(ctx context.Context, tx *sqlx.Tx, productID int64) error {
	_, err := tx.ExecContext(ctx, tx.Rebind(`delete from wp_product_img where product_id = ?`), productID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) updateProductImage(ctx context.Context, tx *sqlx.Tx, productID int64, images []entity.ProductImg) error {

	qImg := `INSERT INTO 
    			wp_product_img (
    			    product_id, 
					image_url, 
    			    created_time, 
    			    created_by, 
    			    updated_time, 
					updated_by
				) VALUES (?, ?, now(), ?, now(), ?)`

	for _, img := range images {
		_, err := tx.ExecContext(ctx, tx.Rebind(qImg), productID, img.ImageURL, img.CreatedBy, img.CreatedBy)
		if err != nil {
			return err
		}
	}

	return nil
}

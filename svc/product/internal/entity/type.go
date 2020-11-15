package entity

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	// ErrIsExist error type for data conflict
	ErrIsExist = errors.New("product is exist")
	// ErrNotFound cluster not found
	ErrNotFound = errors.New("product not found")
)

// Product struct
type Product struct {
	ID          int64     `json:"id" db:"id"`
	ShopID      int64     `json:"shop_id" db:"shop_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Price       int64     `json:"price" db:"price"`
	Status      int       `json:"status" db:"status"`
	CreatedTime time.Time `json:"created_time" db:"created_time"`
	CreatedBy   int64     `json:"created_by" db:"created_by,omitempty"`
	UpdatedTime time.Time `json:"updated_time" db:"updated_time"`
	UpdatedBy   int64     `json:"updated_by" db:"updated_by,omitempty"`
}

// ProductImg struct
type ProductImg struct {
	ProductID   int64     `json:"product_id" db:"product_id"`
	ImageURL    string    `json:"image_url" db:"image_url"`
	CreatedTime time.Time `json:"created_time" db:"created_time"`
	CreatedBy   int64     `json:"created_by" db:"created_by,omitempty"`
	UpdatedTime time.Time `json:"updated_time" db:"updated_time"`
	UpdatedBy   int64     `json:"updated_by" db:"updated_by,omitempty"`
}

// ProductDetail struct
type ProductDetail struct {
	Product
	Images []ProductImg
}

// ParamProduct struct
type ParamSaveProduct struct {
	Product
	Images []ProductImg
}

// ParamGetProduct struct
type ParamGetProduct GetDBParam

// Filter filter struct
type Filter struct {
	Field    string
	Value    interface{}
	Operator string
}

// GetDBParam struct
type GetDBParam struct {
	Search  string
	Filters []Filter
	Limit   int
	Offset  int
	Sort    SortBy
}

// SortBy struct
type SortBy struct {
	Field string
	Asc   bool
}

func (f Filter) getWhereFormula() string {
	switch f.Value.(type) {
	case []string, []int, []int64:
		return fmt.Sprintf("%s IN(:%s) ", f.Field, f.Field)
	default:
		if f.Operator == "" || f.Operator == "=" {
			return fmt.Sprintf("%s = :%s ", f.Field, f.Field)
		}

		return fmt.Sprintf("%s %s :%s ", f.Field, f.Operator, f.Field)
	}
}

// GenerateNamedWhereQueryFormula unified the generate where formula to reduce code duplication
func (g GetDBParam) GenerateNamedWhereQueryFormula() string {
	where := ""

	if g.Search != "" {
		where = fmt.Sprintf(`WHERE LOWER(name) LIKE '%%%s%%' `, strings.ToLower(g.Search))
	}

	for k, v := range g.Filters {
		if k == 0 && !strings.Contains(where, "WHERE") {
			where = "WHERE "
		}

		where += v.getWhereFormula()

		if k == 0 && len(g.Filters) > 1 {
			where += "AND "
		}
	}

	return where
}

// GenerateNamedQueryArgs generate named query from GetDBParam
func (g GetDBParam) GenerateNamedQueryArgs() map[string]interface{} {
	args := map[string]interface{}{
		"limit":  g.Limit,
		"offset": g.Offset,
	}

	for _, v := range g.Filters {
		args[v.Field] = v.Value
	}

	return args
}

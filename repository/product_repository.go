package repository

import (
	"context"
	"database/sql"
	"golang-rest-api/model/domain"
	"golang-rest-api/model/web"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, product domain.Product)
	FindById(ctx context.Context, tx *sql.Tx, productId int) (web.ProductResponse, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []web.ProductResponse
}

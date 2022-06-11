package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-rest-api/helper"
	"golang-rest-api/model/domain"
)

type OrderProductRepositoryImpl struct {
}

func NewOrderProductRepository() *OrderProductRepositoryImpl {
	return &OrderProductRepositoryImpl{}
}

func (o OrderProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, order_product domain.OrderProduct) domain.OrderProduct {
	SQL := "insert into order_roduct(order_id, product_id, qty, price, amount) values (?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, order_product.OrderId, order_product.ProductId, order_product.Qty, order_product.Price, order_product.Amount)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	order_product.Id = int(id)
	return order_product
}

func (o OrderProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, order_product domain.OrderProduct) domain.OrderProduct {
	SQL := "update customer set order_id = ?, product_id = ?, qty = ?, price = ?, amount = ?, where id = ?"
	_, err := tx.ExecContext(ctx, SQL, order_product.OrderId, order_product.ProductId, order_product.Qty, order_product.Price, order_product.Amount, order_product.Id)
	helper.PanicIfError(err)

	return order_product

}

func (o OrderProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, order_product domain.OrderProduct) {
	SQL := "delete from order_product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, order_product.Id)
	helper.PanicIfError(err)
}

func (o OrderProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, order_productId int) (domain.OrderProduct, error) {
	SQL := "select id, order_id, product_id, qty, price, amount from order_product where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, order_productId)
	helper.PanicIfError(err)
	defer rows.Close()

	order_product := domain.OrderProduct{}
	if rows.Next() {
		err := rows.Scan(&order_product.Id, &order_product.OrderId, &order_product.ProductId, &order_product.Qty, &order_product.Price, &order_product.Amount)
		helper.PanicIfError(err)
		return order_product, nil
	} else {
		return order_product, errors.New("order_product is not found")
	}
}

func (o OrderProductRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.OrderProduct {
	SQL := "select id, order_id, product_id, qty, price, amount from order_product"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var order_products []domain.OrderProduct
	for rows.Next() {
		order_product := domain.OrderProduct{}
		err := rows.Scan(&order_product.Id, &order_product.OrderId, &order_product.ProductId, &order_product.Qty, &order_product.Price, &order_product.Amount)
		helper.PanicIfError(err)
		order_products = append(order_products, order_product)
	}
	return order_products
}

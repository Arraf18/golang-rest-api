package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-rest-api/helper"
	"golang-rest-api/model/domain"
)

type OrdersRepositoryImpl struct {
}

func NewOrdersRepository() OrdersRepository {
	return &OrdersRepositoryImpl{}
}

func (o OrdersRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders {
	SQL := "insert into orders(customer_id, total_amount) values (?,?)"
	result, err := tx.ExecContext(ctx, SQL, orders.CustomerId, orders.TotalAmount)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	orders.Id = int(id)
	return orders
}

func (o OrdersRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, orders domain.Orders) domain.Orders {
	SQL := "update orders set customer_id = ?, total_amount = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, orders.CustomerId, orders.TotalAmount, orders.Id)
	helper.PanicIfError(err)

	return orders
}

func (o OrdersRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, orders domain.Orders) {
	SQL := "delete from orders where id = ?"
	_, err := tx.ExecContext(ctx, SQL, orders.Id)
	helper.PanicIfError(err)
}

func (o OrdersRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, ordersId int) (domain.Orders, error) {
	SQL := "select id, customer_id, total_amount from orders where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, ordersId)
	helper.PanicIfError(err)
	defer rows.Close()

	orders := domain.Orders{}
	if rows.Next() {
		err := rows.Scan(&orders.Id, &orders.CustomerId, &orders.TotalAmount)
		helper.PanicIfError(err)
		return orders, nil
	} else {
		return orders, errors.New("orders is not found")
	}
}

func (o OrdersRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.Orders {
	SQL := "select id, customer_id, total_amount from orders"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var orderss []domain.Orders
	for rows.Next() {
		orders := domain.Orders{}
		err := rows.Scan(&orders.Id, &orders.CustomerId, &orders.TotalAmount)
		helper.PanicIfError(err)
		orderss = append(orderss, orders)
	}
	return orderss
}

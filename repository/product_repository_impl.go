package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"golang-rest-api/helper"
	"golang-rest-api/model/domain"
	"golang-rest-api/model/web"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (p ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "insert into product(name, price, category_id) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.CategoryId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	return product
}

func (p ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "update product set name = ?, price = ?, category_id = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.CategoryId, product.Id)
	helper.PanicIfError(err)

	return product
}

func (p ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	SQL := "delete from product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)
}

func (p ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (web.ProductResponse, error) {
	//SELECT product.id, product.name, product.price, category.name as situs_name FROM product
	//INNER JOIN category on product.category_id = category.id;
	logrus.Info("Product Repository Start")
	SQL := "select p.id, p.name, p.price, p.category_id, c.name from product p inner join category c on p.category_id = c.id where p.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	product := web.ProductResponse{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId, &product.CategoryName)
		helper.PanicIfError(err)
		logrus.Info("Product Repository End")
		return product, nil
	} else {
		logrus.Info("Product Repository End")
		return product, errors.New("product is not found")
	}
}

func (p ProductRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []web.ProductResponse {
	//SELECT product.id, product.name, product.price, category.name as situs_name FROM product
	//INNER JOIN category on product.category_id = category.id;
	//SQL := "SELECT product.id, product.name, product.price, product.category_id, category.name as category_name FROM product INNER JOIN category on product.category_id = category.id"
	SQL := "select p.id, p.name, p.price, p.category_id, c.name from product p inner join category c on p.category_id = c.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []web.ProductResponse
	for rows.Next() {
		product := web.ProductResponse{}
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId, &product.CategoryName)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	return products
}

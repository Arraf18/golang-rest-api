package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"golang-rest-api/exception"
	"golang-rest-api/helper"
	"golang-rest-api/model/domain"
	"golang-rest-api/model/web"
	"golang-rest-api/repository"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := domain.Product{
		Name:       request.Name,
		Price:      request.Price,
		CategoryId: request.CategoryId,
		CreatedAt:  request.CreatedAt,
		UpdatedAt:  request.UpdatedAt,
	}

	product = service.ProductRepository.Save(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productResponse, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	product := helper.ToProduct(productResponse)

	product.Name = request.Name
	product.Price = request.Price
	product.CategoryId = request.CategoryId

	product = service.ProductRepository.Update(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service ProductServiceImpl) Delete(ctx context.Context, productId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productResponse, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	product := helper.ToProduct(productResponse)
	service.ProductRepository.Delete(ctx, tx, product)
}

func (service ProductServiceImpl) FindById(ctx context.Context, productId int) web.ProductResponse {
	logrus.Info("Product Service Start")
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	logrus.Info("Product Service End")
	return product
}

func (service ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindByAll(ctx, tx)

	//return helper.ToProductResponses(products)
	return products
}

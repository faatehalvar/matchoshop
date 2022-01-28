package port

import (
	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/dto"
)

type IProductRepo interface {
	Insert(data *domain.Product) (*domain.Product, *errs.AppError)
	CheckByIDAndMerchantID(productID int64, merchantID int64) (bool, *errs.AppError)
	CheckBySKUAndMerchantID(sku string, merchantID int64) (bool, *errs.AppError)
	CheckByIDAndSKUAndMerchantID(productID int64, sku string, merchantID int64) (bool, *errs.AppError)
	GetAllByMerchantID(merchantID int64) ([]domain.ProductList, *errs.AppError)
	GetOneByIDAndMerchantID(productID int64, merchantID int64) (*domain.ProductDetail, *errs.AppError)
	Update(productID int64, data *domain.Product) *errs.AppError
	Delete(productID int64) *errs.AppError
}

type IProductService interface {
	Create(data *dto.CreateProductRequest) (*dto.ResponseData, *errs.AppError)
	GetList(merchantID int64) (*dto.ResponseData, *errs.AppError)
	GetDetail(productID int64, merchantID int64) (*dto.ResponseData, *errs.AppError)
	Update(productID int64, data *dto.CreateProductRequest) (*dto.ResponseData, *errs.AppError)
	Delete(productID int64, merchantID int64) (*dto.ResponseData, *errs.AppError)
}
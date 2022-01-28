package service

import (
	"fmt"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
	"github.com/danisbagus/matchoshop/internal/dto"
)

type ProductService struct {
	repo                       port.IProductRepo
	productCategoryRepo        port.IProductCategoryRepo
	productProductCategoryRepo port.IProductProductCategoryRepo
}

func NewProductService(repo port.IProductRepo, productCategoryRepo port.IProductCategoryRepo, productProductCategoryRepo port.IProductProductCategoryRepo) port.IProductService {
	return &ProductService{
		repo:                       repo,
		productCategoryRepo:        productCategoryRepo,
		productProductCategoryRepo: productProductCategoryRepo,
	}
}

func (r ProductService) Create(req *dto.CreateProductRequest) (*dto.ResponseData, *errs.AppError) {

	appErr := req.Validate()
	if appErr != nil {
		return nil, appErr
	}

	checkProduct, appErr := r.repo.CheckBySKUAndMerchantID(req.Sku, req.MerchantID)
	if appErr != nil {
		return nil, appErr
	}

	if checkProduct {
		errorMessage := fmt.Sprintf("Product with SKU %s is already exits", req.Sku)
		return nil, errs.NewBadRequestError(errorMessage)
	}

	for _, productCategoryID := range req.ProductCategoryID {
		checkProductCategory, appErr := r.productCategoryRepo.CheckByIDAndMerchantID(productCategoryID, req.MerchantID)
		if appErr != nil {
			return nil, appErr
		}

		if !checkProductCategory {
			return nil, errs.NewBadRequestError("Product category not found")
		}
	}

	formProduct := domain.Product{
		MerchantID:  req.MerchantID,
		Name:        req.Name,
		Sku:         req.Sku,
		Description: req.Description,
		Price:       req.Price,
		CreatedAt:   time.Now().Format(dbTSLayout),
		UpdatedAt:   time.Now().Format(dbTSLayout),
	}

	newProductData, appErr := r.repo.Insert(&formProduct)
	if appErr != nil {
		return nil, appErr
	}

	formProductProductCategory := make([]domain.ProductProductCategory, 0)

	for _, productCategoryID := range req.ProductCategoryID {
		formProductProductCategory = append(formProductProductCategory, domain.ProductProductCategory{
			ProductID:         newProductData.ProductID,
			ProductCategoryID: productCategoryID,
		})
	}

	appErr = r.productProductCategoryRepo.BulkInsert(formProductProductCategory)
	if appErr != nil {
		return nil, appErr
	}

	response := dto.NewCreateProductResponse("Sucessfully create data", newProductData)
	return response, nil
}

func (r ProductService) GetList(merchantID int64) (*dto.ResponseData, *errs.AppError) {

	products, appErr := r.repo.GetAllByMerchantID(merchantID)
	if appErr != nil {
		return nil, appErr
	}

	response := dto.NewGetProductListResponse("Successfully get data", products)

	return response, nil
}

func (r ProductService) GetDetail(productID int64, merchantID int64) (*dto.ResponseData, *errs.AppError) {

	product, appErr := r.repo.GetOneByIDAndMerchantID(productID, merchantID)
	if appErr != nil {
		return nil, appErr
	}

	productCategories, appErr := r.productCategoryRepo.GetAllByProductIDAndMerchantID(productID, merchantID)
	if appErr != nil {
		return nil, appErr
	}

	product.ProductCategories = productCategories

	response := dto.NewGetProductDetailResponse("Successfully get data", product)

	return response, nil
}

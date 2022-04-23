package service

import (
	"sync"
	"time"

	"github.com/danisbagus/go-common-packages/errs"
	"github.com/danisbagus/go-common-packages/logger"
	"github.com/danisbagus/matchoshop/internal/core/domain"
	"github.com/danisbagus/matchoshop/internal/core/port"
)

type (
	OrderService struct {
		repo              port.OrderRepo
		repoOrderProduct  port.OrderProductRepo
		repoPaymentResult port.PaymentResultRepo
		repoUser          port.UserRepo
	}
)

func NewOrderService(repo port.OrderRepo, repoOrderProduct port.OrderProductRepo, repoPaymentResult port.PaymentResultRepo, repoUser port.UserRepo) port.OrderService {
	return &OrderService{
		repo:              repo,
		repoOrderProduct:  repoOrderProduct,
		repoPaymentResult: repoPaymentResult,
		repoUser:          repoUser,
	}
}

func (s OrderService) Create(form *domain.OrderDetail) (*domain.OrderDetail, *errs.AppError) {

	form.CreatedAt = time.Now()
	form.UpdatedAt = time.Now()

	orderID, appErr := s.repo.Insert(form)
	if appErr != nil {
		return nil, appErr
	}

	form.Order.OrderID = orderID
	return form, nil
}

func (s OrderService) GetList(userID int64) ([]domain.OrderDetail, *errs.AppError) {
	orders, appErr := s.repo.GetAllByUserID(userID)
	if appErr != nil {
		return nil, appErr
	}
	return orders, nil
}

func (s OrderService) GetDetail(ID int64) (*domain.OrderDetail, *errs.AppError) {

	var order *domain.OrderDetail
	var orderProducts []domain.OrderProduct

	orderChan := make(chan *domain.OrderDetail, 1)
	orderProductChan := make(chan []domain.OrderProduct, 1)
	errorChan := make(chan *errs.AppError, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		order, appErr := s.repo.GetOneByID(ID)
		if appErr != nil {
			errorChan <- appErr
		}
		orderChan <- order
	}()

	go func() {
		defer wg.Done()
		orderProducts, appErr := s.repoOrderProduct.GetAllByOrderID(ID)
		if appErr != nil {
			errorChan <- appErr
		}
		orderProductChan <- orderProducts
	}()

	wg.Wait()

	close(errorChan)
	close(orderChan)
	close(orderProductChan)

	for appErr := range errorChan {
		if appErr != nil {
			return nil, appErr
		}
	}

	for dataChan := range orderChan {
		order = dataChan
	}

	for dataChan := range orderProductChan {
		orderProducts = dataChan
	}

	order.OrderProducts = orderProducts
	return order, nil
}

func (s OrderService) UpdatePaid(form *domain.PaymentResult) *errs.AppError {

	_, appErr := s.repo.GetOneByID(form.OrderID)
	if appErr != nil {
		return appErr
	}

	// check payment result by id
	checkPaymentResult, appErr := s.repoPaymentResult.CheckByID(form.PaymentResultID)
	if appErr != nil {
		return appErr
	}
	if checkPaymentResult {
		logger.Error("Failed while update order paid: payment result id already exist")
		return errs.NewBadRequestError("payment result id already exist")
	}

	// check order payment result
	checkPaymentResult, appErr = s.repoPaymentResult.CheckByOrderIDAndStatus(form.OrderID, "COMPLETED")
	if appErr != nil {
		return appErr
	}

	if checkPaymentResult {
		logger.Error("Failed while update order paid: order already paid")
		return errs.NewBadRequestError("order already paid")
	}

	appErr = s.repo.UpdatePaid(form)
	if appErr != nil {
		return appErr
	}

	return nil
}

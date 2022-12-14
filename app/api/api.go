package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/danisbagus/matchoshop/app/api/middleware"
	"github.com/danisbagus/matchoshop/internal/core/service"
	handlerV1 "github.com/danisbagus/matchoshop/internal/handler/v1"
	"github.com/danisbagus/matchoshop/internal/repo"
	"github.com/danisbagus/matchoshop/utils/config"
	"github.com/danisbagus/matchoshop/utils/constants"

	_ "github.com/lib/pq"
)

func StartApp() {
	config.LoadConfig()

	sqliteDB := config.GetSqliteDB()

	router := mux.NewRouter()

	// wiring
	userRepo := repo.NewUserRepo(sqliteDB)
	productRepo := repo.NewProductRepo(sqliteDB)
	productCategoryRepo := repo.NewProductCategoryRepo(sqliteDB)
	productProductCategoryRepo := repo.NewProductProductCategoryRepo(sqliteDB)
	refreshTokenStoreRepo := repo.NewRefreshTokenStoreRepo(sqliteDB)
	orderRepo := repo.NewOrderRepo(sqliteDB)
	orderProductRepo := repo.NewOrderProductRepo(sqliteDB)
	paymentResultRepo := repo.NewPaymentResult(sqliteDB)
	reviewRepo := repo.NewReviewRepo(sqliteDB)

	userService := service.NewUserService(userRepo, refreshTokenStoreRepo)
	productService := service.NewProductService(productRepo, productCategoryRepo, productProductCategoryRepo, reviewRepo)
	productCategoryService := service.NewProductCategoryService(productCategoryRepo)
	orderService := service.NewOrderService(orderRepo, orderProductRepo, paymentResultRepo, productRepo)
	uploadService := service.NewUploadService()
	reviewService := service.NewReviewService(reviewRepo)

	userHandlerV1 := handlerV1.UserHandler{Service: userService}
	productHandlerV1 := handlerV1.ProductHandler{Service: productService}
	productCategoryHandlerV1 := handlerV1.ProductCategoryHandler{Service: productCategoryService}
	orderHandlerV1 := handlerV1.OrderHandler{Service: orderService}
	configHandlerV1 := handlerV1.ConfigHandler{}
	uploadHandlerV1 := handlerV1.UploadHandler{Service: uploadService}
	reviewHandlerV1 := handlerV1.ReviewHandler{Service: reviewService}

	// auth v1 routes
	authV1Route := router.PathPrefix("/api/v1/auth").Subrouter()
	authV1Route.HandleFunc("/login", userHandlerV1.Login).Methods(http.MethodPost)
	authV1Route.HandleFunc("/refresh", userHandlerV1.Refresh).Methods(http.MethodPost)
	authV1Route.HandleFunc("/register/customer", userHandlerV1.RegisterCustomer).Methods(http.MethodPost)

	// user v1 routes
	userV1Route := router.PathPrefix("/api/v1/user").Subrouter()
	userV1Route.Use(middleware.AuthorizationHandler())
	userV1Route.HandleFunc("", userHandlerV1.GetUserDetail).Methods(http.MethodGet)
	userV1Route.HandleFunc("/profile", userHandlerV1.UpdateUser).Methods(http.MethodPatch)

	// product v1 routes
	productV1Route := router.PathPrefix("/api/v1/product").Subrouter()
	productV1Route.HandleFunc("", productHandlerV1.GetProductListPaginate).Methods(http.MethodGet)
	productV1Route.HandleFunc("/top", productHandlerV1.GetTopProduct).Methods(http.MethodGet)
	productV1Route.HandleFunc("/{product_id}", productHandlerV1.GetProductDetail).Methods(http.MethodGet)

	// product admin v1 routes
	productAdminV1Route := router.PathPrefix("/api/v1/admin/product").Subrouter()
	productAdminV1Route.Use(middleware.AuthorizationHandler(), middleware.ACL(constants.AdminPermission))
	productAdminV1Route.HandleFunc("", productHandlerV1.CreateProduct).Methods(http.MethodPost)
	productAdminV1Route.HandleFunc("", productHandlerV1.GetProductListPaginate).Methods(http.MethodGet)
	productAdminV1Route.HandleFunc("/{product_id}", productHandlerV1.UpdateProduct).Methods(http.MethodPut)
	productAdminV1Route.HandleFunc("/{product_id}", productHandlerV1.Delete).Methods(http.MethodDelete)
	productAdminV1Route.HandleFunc("/{product_id}", productHandlerV1.GetProductDetail).Methods(http.MethodGet)

	// product category v1 routes
	productCategoryV1Route := router.PathPrefix("/api/v1/product-category").Subrouter()
	productCategoryV1Route.HandleFunc("", productCategoryHandlerV1.GetProductCategoryList).Methods(http.MethodGet)
	productCategoryV1Route.HandleFunc("/{product_category_id}", productCategoryHandlerV1.GetProductCategoryDetail).Methods(http.MethodGet)

	// product category admin v1 routes
	productCategoryAdminV1Route := router.PathPrefix("/api/v1/admin/product-category").Subrouter()
	productAdminV1Route.Use(middleware.AuthorizationHandler(), middleware.ACL(constants.AdminPermission))
	productCategoryAdminV1Route.HandleFunc("", productCategoryHandlerV1.CreateProductCategory).Methods(http.MethodPost)
	productCategoryAdminV1Route.HandleFunc("/{product_category_id}", productCategoryHandlerV1.UpdateProductCategory).Methods(http.MethodPut)
	productCategoryAdminV1Route.HandleFunc("/{product_category_id}", productCategoryHandlerV1.Delete).Methods(http.MethodDelete)

	// order v1 routes
	orderV1Route := router.PathPrefix("/api/v1/order").Subrouter()
	orderV1Route.Use(middleware.AuthorizationHandler(), middleware.ACL(constants.CustomerPermission))
	orderV1Route.HandleFunc("", orderHandlerV1.Create).Methods(http.MethodPost)
	orderV1Route.HandleFunc("", orderHandlerV1.GetList).Methods(http.MethodGet)
	orderV1Route.HandleFunc("/{order_id}", orderHandlerV1.GetDetail).Methods(http.MethodGet)
	orderV1Route.HandleFunc("/{order_id}/pay", orderHandlerV1.UpdatePaid).Methods(http.MethodPut)

	// product admin v1 routes
	userAdminV1Route := router.PathPrefix("/api/v1/admin/user").Subrouter()
	userAdminV1Route.Use(middleware.AuthorizationHandler(), middleware.ACL(constants.AdminPermission))
	userAdminV1Route.HandleFunc("", userHandlerV1.GetUserList).Methods(http.MethodGet)
	userAdminV1Route.HandleFunc("/{user_id}", userHandlerV1.GetUserDetailAdmin).Methods(http.MethodGet)
	userAdminV1Route.HandleFunc("/{user_id}", userHandlerV1.DeleteUser).Methods(http.MethodDelete)
	userAdminV1Route.HandleFunc("/{user_id}", userHandlerV1.UpdateUserAdmin).Methods(http.MethodPatch)

	// config v1 routes
	configRoute := router.PathPrefix("/api/v1/config").Subrouter()
	configRoute.HandleFunc("/paypal", configHandlerV1.GetPaypalConfig).Methods(http.MethodGet)

	// admin config v1 routes
	adminConfigRoute := router.PathPrefix("/api/v1/admin/config").Subrouter()
	adminConfigRoute.HandleFunc("", configHandlerV1.GetConfig).Methods(http.MethodGet)

	// upload v1 routes
	uploadRoute := router.PathPrefix("/api/v1/upload").Subrouter()
	uploadRoute.HandleFunc("/image", uploadHandlerV1.UploadImage).Methods(http.MethodPost)

	// order admin v1 routes
	orderAdminV1Route := router.PathPrefix("/api/v1/admin/order").Subrouter()
	orderAdminV1Route.Use(middleware.AuthorizationHandler(), middleware.ACL(constants.AdminPermission))
	orderAdminV1Route.HandleFunc("", orderHandlerV1.GetListAdmin).Methods(http.MethodGet)
	orderAdminV1Route.HandleFunc("/{order_id}", orderHandlerV1.GetDetail).Methods(http.MethodGet)
	orderAdminV1Route.HandleFunc("/{order_id}/deliver", orderHandlerV1.UpdateDelivered).Methods(http.MethodPut)

	// review v1 routes
	reviewV1Route := router.PathPrefix("/api/v1/review").Subrouter()
	reviewV1Route.Use(middleware.AuthorizationHandler(), middleware.ACL(constants.CustomerPermission))
	reviewV1Route.HandleFunc("", reviewHandlerV1.Create).Methods(http.MethodPost)
	reviewV1Route.HandleFunc("", reviewHandlerV1.Update).Methods(http.MethodPut)
	reviewV1Route.HandleFunc("/{product_id}", reviewHandlerV1.GetDetail).Methods(http.MethodGet)

	router.HandleFunc("/health-check", healthCheck)

	port := config.GetPort()
	addr := fmt.Sprintf(":%v", port)
	fmt.Println("Starting the application at:", port)
	log.Fatal(http.ListenAndServe(addr, router))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("App Up"))
}

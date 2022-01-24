package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	"github.com/danisbagus/matchoshop/internal/core/service"
	handlerV1 "github.com/danisbagus/matchoshop/internal/handler/v1"
	"github.com/danisbagus/matchoshop/internal/repo"

	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := GetClient()
	defer client.Close()

	router := mux.NewRouter()

	userRepo := repo.NewUserRepo(client)

	userService := service.NewUserService(userRepo)

	userHandlerV1 := handlerV1.AuthHandler{Service: userService}

	authRouter := router.PathPrefix("/auth").Subrouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	// v1 route
	authRouter.HandleFunc("/v1/login", userHandlerV1.Login).Methods(http.MethodPost)
	authRouter.HandleFunc("/v1/register/merchant", userHandlerV1.RegisterMerchant).Methods(http.MethodPost)

	apiRouter.HandleFunc("/hello", SayHello)

	APP_PORT := os.Getenv("APP_PORT")
	if APP_PORT == "" {
		log.Fatal("$APP_PORT must be set")
	}

	appAPP_PORT := fmt.Sprintf("%v:%v", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	fmt.Println("Starting the application at:", appAPP_PORT)
	log.Fatal(http.ListenAndServe(appAPP_PORT, router))
}

func GetClient() *sqlx.DB {
	dbAPP_HOST := os.Getenv("DB_APP_HOST")
	dbAPP_PORT := os.Getenv("DB_APP_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connection := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPassword,
		dbAPP_HOST,
		dbAPP_PORT,
		dbName,
	)

	client, err := sqlx.Open("postgres", connection)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}

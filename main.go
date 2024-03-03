package main

import (
	"context"
	"fmt"
	"github.com/dreis0/rinha-backend/domain/usecases"
	"github.com/dreis0/rinha-backend/gateways/postgres"
	"github.com/dreis0/rinha-backend/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	router := mux.NewRouter()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		panic("PORT environment variable is not set")
	}

	s := &http.Server{
		Addr:         fmt.Sprintf("127.0.0.1:%s", port),
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	conn := initializePostgresConnection()
	dal := postgres.NewPostgresDal(*conn)
	usecase := usecases.New(dal)

	tansactionsHandler := handlers.NewTransactionsHandler(usecase)
	statementHandler := handlers.NewStatementHandler(usecase)

	router.HandleFunc("/clientes/{id:[0-9]+}/extrato", statementHandler.GetCustomerStatement).Methods("GET")
	router.HandleFunc("/clientes/{id:[0-9]+}/transacoes", tansactionsHandler.DoTransaction).Methods("POST")

	log.Printf("Server running on port %s", port)
	log.Fatal(s.ListenAndServe())
}

func initializePostgresConnection() *pgx.Conn {
	db, ok := os.LookupEnv("POSTGRES_DATABASE")
	if !ok {
		panic("POSTGRES_DATABASE environment variable is not set")
	}

	hostname, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		panic("POSTGRES_HOST environment variable is not set")
	}

	port, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		panic("POSTGRES_PORT environment variable is not set")
	}

	user, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		panic("POSTGRES_USER environment variable is not set")
	}

	password, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		panic("POSTGRES_PASSWORD environment variable is not set")
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, hostname, port, db)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	return conn
}

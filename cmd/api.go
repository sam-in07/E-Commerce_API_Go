package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

type application struct {
	config config
	// logger
	db *pgx.Conn
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID) // important for rate limiting
	r.Use(middleware.RealIP)    // import for rate limiting and analytics and tracing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from crashes

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})

	// productService := products.NewService(repo.New(app.db))
	// productHandler := products.NewHandler(productService)
	// r.Get("/products", productHandler.ListProducts)

	// orderService := orders.NewService(repo.New(app.db), app.db)
	// ordersHandler := orders.NewHandler(orderService)
	// r.Post("/orders", ordersHandler.PlaceOrder)

	return r
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

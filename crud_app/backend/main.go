package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"

	"crud_app/config"
	"crud_app/db"
	"crud_app/handlers"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	usersHandler := &handlers.UsersHandler{
		DB: db,
	}

	r := mux.NewRouter()
	r.HandleFunc("/users", usersHandler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users", usersHandler.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{id:[0-9]+}", usersHandler.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/users/{id:[0-9]+}", usersHandler.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/users/{id:[0-9]+}", usersHandler.DeleteUser).Methods(http.MethodDelete)

	// Agregar tus rutas a tu manejador HTTP

	// Crear opciones de cors
	corsOptions := cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Agrega los dominios permitidos
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}

	// Agregar cors a tu manejador HTTP
	c := cors.New(corsOptions)
	handlerWithCors := c.Handler(r)

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Listening on %s", addr)
	err = http.ListenAndServe(addr, handlerWithCors)
	if err != nil {
		log.Fatal(err)
	}
}

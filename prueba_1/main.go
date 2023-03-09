package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Order struct {
	OrderID int    `json:"order_id"`
	Date    string `json:"date"`
	Client  struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Address struct {
			Street     string `json:"street"`
			Number     int    `json:"number"`
			City       string `json:"city"`
			Country    string `json:"country"`
			PostalCode int    `json:"postal_code"`
		} `json:"address"`
	} `json:"client"`
	Items []struct {
		Product  string  `json:"product"`
		Detail   string  `json:"detail"`
		Quantity int     `json:"quantity"`
		Price    float64 `json:"price"`
	} `json:"items"`
	Total float64 `json:"total"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// Decodificar el body del request
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser := User{
		Name:  user.Name,
		Email: user.Email,
	}
	userJSON, err := json.Marshal(newUser)
	// Crear el usuario en la base de datos o hacer cualquier otra tarea
	fmt.Printf("Se creó el usuario con nombre %s e email %s\n", user.Name, user.Email)

	// Escribir la respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(userJSON)
}

func getUser(w http.ResponseWriter, r *http.Request) {

	user := User{
		Name:  "Juan",
		Email: "juan@example.com",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func getOrder(w http.ResponseWriter, r *http.Request) {
	order := Order{
		OrderID: 1234,
		Date:    "2023-03-07",
		Client: struct {
			Name    string `json:"name"`
			Email   string `json:"email"`
			Phone   string `json:"phone"`
			Address struct {
				Street     string `json:"street"`
				Number     int    `json:"number"`
				City       string `json:"city"`
				Country    string `json:"country"`
				PostalCode int    `json:"postal_code"`
			} `json:"address"`
		}{
			Name:  "Jhon Dee",
			Email: "jhondee@gmail.com",
			Phone: "555-1234",
			Address: struct {
				Street     string `json:"street"`
				Number     int    `json:"number"`
				City       string `json:"city"`
				Country    string `json:"country"`
				PostalCode int    `json:"postal_code"`
			}{
				Street:     "New Street",
				Number:     123,
				City:       "Salt Lake City",
				Country:    "United States of America",
				PostalCode: 5016,
			},
		},
		Items: []struct {
			Product  string  `json:"product"`
			Detail   string  `json:"detail"`
			Quantity int     `json:"quantity"`
			Price    float64 `json:"price"`
		}{
			{
				Product:  "Shirt",
				Detail:   "This is a product detail",
				Quantity: 2,
				Price:    25.99,
			},
			{
				Product:  "Pants",
				Detail:   "This is a product detail",
				Quantity: 1,
				Price:    35.99,
			},
		},
		Total: 87.97,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func orderDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Hacer una solicitud a /api/get-order para obtener la información del pedido
	resp, err := http.Get("http://localhost:8080/api/get-order")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Leer la respuesta y decodificarla en una estructura de pedido
	var order Order
	err = json.NewDecoder(resp.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Cargar la plantilla HTML
	tmpl, err := template.ParseFiles("order.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar la plantilla con la información del pedido
	err = tmpl.Execute(w, order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/api/get-user", getUser)
	http.HandleFunc("/api/get-order", getOrder)
	http.HandleFunc("/order-detail", orderDetailHandler)
	http.HandleFunc("/create-user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createUser(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}

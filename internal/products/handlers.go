package products

import (
	"github.com/sam-in07/E-Commerce_API_Go/internal/json"
	"log"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// call the service => ListProduct . 2nd : Return Json in an HTTP response

	err := h.service.ListProducts(r.Context())
	// products := []string{"Hlw", "WERSDS"}
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products := struct {
		Products []string `json:"Products"`
	}{}

	json.Write(w, http.StatusOK, products)
}

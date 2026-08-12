package products

import (
	"encoding/json"
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

	// products, err := h.service.ListProducts(r.Context())
	products := []string{"Hlw", "WERSDS"}
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	
	json.Write(w, http.StatusOK, products)
}

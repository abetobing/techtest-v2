package customer

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"customer/api/middlewares"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	var req CustomerDetail
	json.NewDecoder(r.Body).Decode(&req)
	id, err := CreateCustomer(middlewares.GetDB(r.Context()), req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(CreateResponse{CustomerID: id})
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	res, err := GetCustomer(middlewares.GetDB(r.Context()), id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{Error: http.StatusText(http.StatusNotFound)})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	res, err := ListCustomer(middlewares.GetDB(r.Context()))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	err := DeleteCustomer(middlewares.GetDB(r.Context()), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}

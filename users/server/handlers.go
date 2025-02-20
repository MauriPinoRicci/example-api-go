package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MauriPinoRicci/example-api-go/users/application/create_srv"
	"github.com/MauriPinoRicci/example-api-go/users/application/get_srv"
	"github.com/MauriPinoRicci/example-api-go/users/infra/users_dynamo"
	"github.com/go-chi/chi/v5"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	var input create_srv.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createSrv := create_srv.NewService(users_dynamo.New())

	resp, err := createSrv.Execute(r.Context(), &input)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respByte, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respByte)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}

	getSrv := get_srv.NewService(users_dynamo.New())

	resp, err := getSrv.GetByID(r.Context(), &get_srv.GetUserInput{ID: id})
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respByte, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respByte)
}

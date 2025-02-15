package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MauriPinoRicci/example-api-go/server/application/create_srv"
	"github.com/MauriPinoRicci/example-api-go/server/infra/users_dynamo"
	"github.com/go-chi/chi"
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

	resp, err := createSrv.CreateUser(r.Context(), &input)
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

func getUsers(w http.ResponseWriter, r *http.Request) {
	Id := chi.URLParam(r, "id")

	fmt.Println("ID recido:", Id)

	if Id == "" {
		http.Error(w, "missing user ID", http.StatusBadRequest)
		return
	}

	createSrv := create_srv.NewService(users_dynamo.New())

	resp, err := createSrv.GetUserByID(r.Context(), Id)
	if err != nil {
		fmt.Println("Error fetching user:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respByte, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Error marshalling response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respByte)
}

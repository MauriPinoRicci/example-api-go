package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MauriPinoRicci/example-api-go/users/application/create_srv"
	"github.com/MauriPinoRicci/example-api-go/users/infra/users_dynamo"
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

package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MauriPinoRicci/example-api-go/users/application/create_srv"
	"github.com/MauriPinoRicci/example-api-go/users/application/delete_srv"
	"github.com/MauriPinoRicci/example-api-go/users/application/get_srv"
	"github.com/MauriPinoRicci/example-api-go/users/application/update_srv"
	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	createSrv *create_srv.Service
	getSrv    *get_srv.Service
	updateSrv *update_srv.Service
	deleteSrv *delete_srv.Service
}

func NewHandlers(dep *Dependencies) *Handlers {
	return &Handlers{
		createSrv: dep.createSrv,
		getSrv:    dep.getSrv,
		updateSrv: dep.updateSrv,
		deleteSrv: dep.deleteSrv,
	}
}

func (s *Handlers) create(w http.ResponseWriter, r *http.Request) {
	var input create_srv.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := s.createSrv.Execute(r.Context(), &input)
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

func (s *Handlers) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}

	resp, err := s.getSrv.GetByID(r.Context(), &get_srv.GetUserInput{ID: id})
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

func (s *Handlers) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}

	var input update_srv.UpdateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	input.ID = id
	resp, err := s.updateSrv.Update(r.Context(), &input)
	if err != nil {
		if err.Error() == "user not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
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

func (s *Handlers) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "missing id query parameter", http.StatusBadRequest)
		return
	}

	err := s.deleteSrv.Execute(r.Context(), &delete_srv.DeleteUserInput{ID: id})
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNoContent)
}

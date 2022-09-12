package controllers

import (
	"chi-boilerplate/models"
	"chi-boilerplate/repository"
	"encoding/json"
	"net/http"
	"strconv"
)

type EmpHandler struct {
	repo repository.ExampleRepo
}

func NewHandler(repo repository.ExampleRepo) *EmpHandler {
	return &EmpHandler{
		repo: repo,
	}
}
func (h *EmpHandler) GetData(w http.ResponseWriter, request *http.Request) {
	q := request.URL.Query()
	limit, _ := strconv.Atoi(q.Get("limit"))
	offset, _ := strconv.Atoi(q.Get("offset"))

	data, err := h.repo.GetExamples(int64(limit), int64(offset))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&data)
}
func (h *EmpHandler) CreateData(w http.ResponseWriter, request *http.Request) {
	example := new(models.Example)
	err := json.NewDecoder(request.Body).Decode(&example)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.repo.CreateExample(example)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(example)
}

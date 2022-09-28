package controllers

import (
	"encoding/json"
	"go-fication/helpers"
	"go-fication/infra/logger"
	"go-fication/models"
	"go-fication/repository"
	"net/http"
	"strconv"
)

type ExampleHandler struct {
	repo repository.ExampleRepo
}

func NewExampleHandler(repo repository.ExampleRepo) *ExampleHandler {
	return &ExampleHandler{
		repo: repo,
	}
}
func (h *ExampleHandler) GetData(w http.ResponseWriter, request *http.Request) {
	q := request.URL.Query()
	limit, _ := strconv.Atoi(q.Get("limit"))
	offset, _ := strconv.Atoi(q.Get("offset"))

	data, err := h.repo.GetExamples(int64(limit), int64(offset))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	logger.DebugJson("data", data)
	helpers.HttpResponse(data, w)

}

func (h *ExampleHandler) CreateData(w http.ResponseWriter, request *http.Request) {
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

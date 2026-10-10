package done

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"main.go/core/apperrors"
)

type DoneHandler struct {
	service *Service
}

func NewHandler(service *Service) *DoneHandler {
	return &DoneHandler{
		service: service,
	}
}

func (h DoneHandler) DoneHandler(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := h.service.DoneTask(r.Context(), id)

	if errors.Is(err, apperrors.ErrInvalidTaskID) {
		http.Error(w, "Неккоректный ID задачи", http.StatusBadRequest)
		return
	}

	if errors.Is(err, apperrors.ErrTaskNotFound) {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Ошибка при выполнении задачи", http.StatusInternalServerError)
		return
	}

	response := ToTaskDTO(*task)

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

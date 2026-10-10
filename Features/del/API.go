package del

import (
	"errors"
	"net/http"
	"strconv"

	"main.go/core/apperrors"
)

type DelHandler struct {
	service *Service
}

func NewHandler(service *Service) *DelHandler {
	return &DelHandler{
		service: service,
	}
}

func (h *DelHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")

	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteTask(r.Context(), id)

	if errors.Is(err, apperrors.ErrInvalidTaskID) {
		http.Error(w, "Неккоректный ID задачи", http.StatusBadRequest)
		return
	}

	if errors.Is(err, apperrors.ErrTaskNotFound) {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Ошибка при удалении задачи", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

package list

import (
	"encoding/json"
	"net/http"
)

type ListHandler struct {
	service *Service
}

func NewHandler(service *Service) *ListHandler {
	return &ListHandler{
		service: service,
	}
}

func (h *ListHandler) ListHandler(w http.ResponseWriter, r *http.Request) {

	tasks, err := h.service.TasksList(r.Context())
	if err != nil {
		http.Error(w, "ошибка при получении списка задач", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		return
	}
}

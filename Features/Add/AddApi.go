package Add

import (
	"encoding/json"
	"net/http"
)

type AddHandler struct {
	service *Service
}

func NewHandler(service *Service) *AddHandler {
	return &AddHandler{
		service: service,
	}
}
func (h AddHandler) AddHandler(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO

	err := json.NewDecoder(r.Body).Decode(&taskDTO)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	task, err := h.service.CreateTaskModel(r.Context(), &taskDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

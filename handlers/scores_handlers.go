package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	models "../models"
	repository "../repository"
	utils "../utils"
	"github.com/go-chi/chi"
)

// ScoresHandler ...
type ScoresHandler struct {
	repository repository.ScoresInterface
}

// NewScoresHandler ..
func NewScoresHandler(db *sql.DB) *ScoresHandler {
	return &ScoresHandler{
		repository: repository.NewSQLScores(db),
	}
}

// GetAll ..
func (uh *ScoresHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	payload, err := uh.repository.Fetch(r.Context(), 10)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	utils.RespondwithJSON(w, http.StatusOK, payload)
}

// Get ..
func (uh *ScoresHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	payload, err := uh.repository.GetByID(r.Context(), int64(id))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	utils.RespondwithJSON(w, http.StatusFound, payload)
}

// Create ..
func (uh *ScoresHandler) Create(w http.ResponseWriter, r *http.Request) {
	score := models.Scores{}
	json.NewDecoder(r.Body).Decode(&score)

	newID, err := uh.repository.Create(r.Context(), &score)
	if err != nil {
		fmt.Printf("Create %s\n", err)
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	utils.RespondwithJSON(w, http.StatusCreated, map[string]string{"id": newID})
}

// Update ..
func (uh *ScoresHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	score := models.Scores{}
	json.NewDecoder(r.Body).Decode(&score)
	score.ID = string(id)

	payload, err := uh.repository.Update(r.Context(), &score)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	utils.RespondwithJSON(w, http.StatusOK, payload)
}

// Delete ..
func (uh *ScoresHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := uh.repository.Delete(r.Context(), int64(id))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	utils.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Delete Successfully"})
}

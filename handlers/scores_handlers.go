package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

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

// GetAll is a method to handler request to get scores
func (sh *ScoresHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	lastRn := utils.ParseInt64(r.URL.Query().Get("lastRn"))
	limit := utils.ParseInt64(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 10
	}
	payload, err := sh.repository.Fetch(r.Context(), lastRn, limit)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	utils.RespondwithJSON(w, http.StatusOK, payload)
	return
}

// Get is a method to handler request to get a single score
func (sh *ScoresHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	payload, err := sh.repository.GetByID(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	utils.RespondwithJSON(w, http.StatusFound, payload)
	return
}

// Create is a method to handler request to create a new score
func (sh *ScoresHandler) Create(w http.ResponseWriter, r *http.Request) {
	score := models.Scores{}
	json.NewDecoder(r.Body).Decode(&score)

	newID, err := sh.repository.Create(r.Context(), &score)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	utils.RespondwithJSON(w, http.StatusCreated, map[string]string{"id": newID})
	return
}

// Update is a method to handler request to update a score
func (sh *ScoresHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	score := models.Scores{}
	json.NewDecoder(r.Body).Decode(&score)
	score.ID = id

	payload, err := sh.repository.Update(r.Context(), &score)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	utils.RespondwithJSON(w, http.StatusOK, payload)
	return
}

// Delete is a method to handler request to delete a score
func (sh *ScoresHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	_, errDel := sh.repository.Delete(r.Context(), id)
	if errDel != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	utils.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Delete Successfully"})
	return
}

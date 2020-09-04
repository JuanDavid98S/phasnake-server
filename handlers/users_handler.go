package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	models "../models"
	repository "../repository"
	utils "../utils"

	"github.com/go-chi/chi"
)

// UsersHandler ...
type UsersHandler struct {
	repository repository.UserRepository
}

func NewUserHandler(db *sql.DB) *UsersHandler {
	return &UsersHandler{
		repository: repository.NewSQLUserRepository(db),
	}
}

func (uh *UsersHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	payload, err := uh.repository.Fetch(r.Context(), 5)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	utils.RespondwithJSON(w, http.StatusOK, payload)
}

func (uh *UsersHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	payload, err := uh.repository.GetByID(r.Context(), int64(id))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	utils.RespondwithJSON(w, http.StatusFound, payload)
}

func (uh *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)

	newID, err := uh.repository.Create(r.Context(), &user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	utils.RespondwithJSON(w, http.StatusCreated, map[string]int64{"id": newID})
}

func (uh *UsersHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = int64(id)

	payload, err := uh.repository.Update(r.Context(), &user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	utils.RespondwithJSON(w, http.StatusOK, payload)
}

func (uh *UsersHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := uh.repository.Delete(r.Context(), int64(id))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	utils.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Delete Successfully"})
}

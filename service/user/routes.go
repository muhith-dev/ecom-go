package user

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/muhith-dev/ecom-go/service/auth"
	"github.com/muhith-dev/ecom-go/types"
	"github.com/muhith-dev/ecom-go/utils"
	"net/http"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WeriteError(w, http.StatusBadRequest, err)
	}

	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WeriteError(w, http.StatusBadRequest, fmt.Errorf("user already exists", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)

	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WeriteError(w, http.StatusInternalServerError, err)
		return
	}

}

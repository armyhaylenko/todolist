package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/armyhaylenko/todolist/internal/logging"
	"github.com/armyhaylenko/todolist/internal/models"
	"github.com/redis/go-redis/v9"
)

type authHandler struct {
	redis redis.Client
}

func newAuthHandler(redis redis.Client) authHandler {
	return authHandler{
		redis,
	}
}

func (h *authHandler) authRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", h.register)
	mux.HandleFunc("POST /login", h.login)
	mux.HandleFunc("POST /logout", h.logout)

	return mux
}

func (h *authHandler) register(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	var registerPayload Register

	if err := json.NewDecoder(req.Body).Decode(&registerPayload); err != nil {
		logging.Logger.Errorw("Could not decode register payload!", "error", err)
		http.Error(w, "Could not decode register payload", http.StatusBadRequest)
		return
	}

	key := fmt.Sprint("user:", registerPayload.Email)

	exists, err := h.redis.Exists(ctx, key).Result()
	if err != nil {
		logging.Logger.Errorw("Failed to check if user exists in redis", "error", err, "email", registerPayload.Email)
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}
	if exists != 0 {
		logging.Logger.Errorw("User already exists", "email", registerPayload.Email)
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	user, err := models.New(registerPayload.Email, registerPayload.Password)

	if err != nil {
		logging.Logger.Errorw("Failed to create user from registerPayload", "error", err)
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	if err := h.redis.HSet(ctx, key, user).Err(); err != nil {
		logging.Logger.Errorw("Failed writing user to redis", "error", err, "email", registerPayload.Email)
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *authHandler) login(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	var loginPayload Login

	if err := json.NewDecoder(req.Body).Decode(&loginPayload); err != nil {
		logging.Logger.Errorw("Could not decode login payload!", "error", err)
		http.Error(w, "Could not decode login payload", http.StatusBadRequest)
		return
	}

	key := fmt.Sprint("user:", loginPayload.Email)

	user := &models.User{}
	resp, err := h.redis.HGetAll(ctx, key).Result()
	if err != nil {
		logging.Logger.Errorw("Failed to get user from redis", "error", err, "email", loginPayload.Email)
		http.Error(w, "Login failed", http.StatusInternalServerError)
		return
	}

	if err := user.FromMap(resp); err != nil {
		if err == models.ErrUserNotFound {
			logging.Logger.Errorw("User not found", "error", err, "email", loginPayload.Email)
			http.Error(w, "Invalid username or password", http.StatusBadRequest)
			return
		}
		logging.Logger.Errorw("Failed to deserialize user from redis response", "error", err, "email", loginPayload.Email)
		http.Error(w, "Login failed", http.StatusInternalServerError)
		return
	}
}

func (*authHandler) logout(w http.ResponseWriter, req *http.Request) {
	// TODO: add deletion of access token id from db

}

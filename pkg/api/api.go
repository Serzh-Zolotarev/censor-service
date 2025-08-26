package api

import (
	"censor-service/pkg/censor"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Comment struct {
	Content string `json:"content"`
}

// Программный интерфейс сервера
type API struct {
	router    *mux.Router
	validator func(content string) bool
}

// Конструктор объекта API
func New(validator censor.ValidatorFunc) *API {
	api := API{
		router:    mux.NewRouter(),
		validator: validator,
	}
	api.endpoints()
	return &api
}

// Регистрация обработчиков API
func (a *API) endpoints() {
	a.router.Use(requestIdMiddleware)
	// валидация комментария
	a.router.HandleFunc("/", a.validateHandler).Methods(http.MethodPost, http.MethodOptions)
	a.router.Use(loggingMiddleware)
}

// Получение маршрутизатора запросов
func (a *API) Router() *mux.Router {
	return a.router
}

// Получение всех комментариев к новости
func (a *API) validateHandler(w http.ResponseWriter, r *http.Request) {
	var p Comment
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !a.validator(p.Content) {
		http.Error(w, "comment censored", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

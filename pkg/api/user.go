package api

import (
	"encoding/json"
	"net/http"

	"github.com/tonymontanapaffpaff/golang-training-university/pkg/api/middleware"
	"github.com/tonymontanapaffpaff/golang-training-university/pkg/data"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type userAPI struct {
	data *data.UserData
}

func NewUserAPI(userData *data.UserData) *userAPI {
	return &userAPI{data: userData}
}

func ServeUserResource(r *mux.Router, data data.UserData) {
	api := &userAPI{data: &data}
	r.HandleFunc("/login", api.Login).Methods("POST")
	r.HandleFunc("/refresh", api.Refresh).Methods("POST")
	subRouter := r.Methods("POST").Subrouter()
	subRouter.HandleFunc("/logout", api.Logout)
	subRouter.Use(middleware.TokenAuthMiddleware)
}

func (a *userAPI) Login(writer http.ResponseWriter, request *http.Request) {
	var user data.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Errorf("failed reading JSON: %s", err)
		writer.WriteHeader(http.StatusUnprocessableEntity)
		_, err = writer.Write([]byte("Invalid json provided"))
		if err != nil {
			log.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	tokens, err := a.data.Login(user)
	if err != nil {
		log.Error(err)
		return
	}
	jsonString, err := json.Marshal(tokens)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(jsonString)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a *userAPI) Logout(writer http.ResponseWriter, request *http.Request) {
	err := a.data.Logout(request)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		_, err := writer.Write([]byte(err.Error()))
		if err != nil {
			log.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write([]byte("Successfully logged out"))
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a *userAPI) Refresh(writer http.ResponseWriter, request *http.Request) {
	mapToken := make(map[string]string)
	err := json.NewDecoder(request.Body).Decode(&mapToken)
	if err != nil {
		log.Errorf("failed reading JSON: %s", err)
		writer.WriteHeader(http.StatusUnprocessableEntity)
		_, err = writer.Write([]byte("Invalid json provided"))
		if err != nil {
			log.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	refreshToken := mapToken["refresh_token"]
	newPair, refreshErr := a.data.Refresh(refreshToken)
	if refreshErr.Err != nil {
		writer.WriteHeader(refreshErr.StatusCode)
		_, err := writer.Write([]byte(refreshErr.Err.Error()))
		if err != nil {
			log.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	jsonString, err := json.Marshal(newPair)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	_, err = writer.Write(jsonString)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

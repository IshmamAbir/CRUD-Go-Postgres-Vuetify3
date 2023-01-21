package http

import (
	"backend/models"
	"backend/user"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type userHandler struct {
	UserUsecase user.Usecase
}

func NewUserHandler(router *mux.Router, uu user.Usecase) {
	handler := &userHandler{
		UserUsecase: uu,
	}

	//mapping to the url
	router.HandleFunc("/users", handler.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handler.GetUserById).Methods("GET")
	router.HandleFunc("/users", handler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handler.deleteUser).Methods("DELETE")
}

//----------------------------------------------------------------------------------------

// Get All Users method
func (h *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := h.UserUsecase.GetAllUsers()
	if err != nil {
		log.Panic(err)
		json.NewEncoder(w).Encode(err.Error())
	} else if users == nil {
		json.NewEncoder(w).Encode("No User Found!")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// Get user by Id
func (h *userHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// taking route values from the current request
	vars := mux.Vars(r)
	// taking uuid from the route path variable
	userId, err := uuid.Parse(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Error parsing Id")
		return
	}
	user, err1 := h.UserUsecase.GetUserById(userId)
	if err1 != nil {
		json.NewEncoder(w).Encode(err1.Error())
		return
	}
	json.NewEncoder(w).Encode(user)
}

// create user
func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	inputErr := h.UserUsecase.SaveUser(&user)
	if inputErr != nil {
		json.NewEncoder(w).Encode(inputErr.Error())
		return
	}
	json.NewEncoder(w).Encode(user)
}

// update user
func (h userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	vars := mux.Vars(r)
	userId, err := uuid.Parse(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Error parsing Id")
		return
	}
	user := models.User{}
	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode("Error parsing the input data")
		return
	}
	updateErr := h.UserUsecase.UpdateUser(userId, &user)
	if updateErr != nil {
		json.NewEncoder(w).Encode(updateErr.Error())
		return
	}
	json.NewEncoder(w).Encode(user)
}

// delete user
func (h userHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	vars := mux.Vars(r)
	userId, err := uuid.Parse(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Error parsing Id")
		return
	}
	if err := h.UserUsecase.DeleteUer(userId); err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("Delete Successful")
}

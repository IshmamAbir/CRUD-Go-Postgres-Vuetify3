package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// module name is 'backend'. check go.mod for your package module name
	_Util "backend/util"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

var db *gorm.DB
var err error

func main() {
	fmt.Println("Running")
	fmt.Println("--------------------------")

	//step 1: First initialize the database connection.
	// I used gorm for database connection
	dbInit()

	// step 2: Set the routing url path with their method type and methods definition
	// I used mux routing for the url path declaration
	routers()
}

func routers() {
	router := mux.NewRouter()
	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUer).Methods("DELETE")

	http.ListenAndServe(":9080", &CORSRouterDecorator{router})

}

//-----------------------------------------------------------------------------

type CORSRouterDecorator struct {
	R *mux.Router
}

func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter,
	req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Accept-Language,"+
				" Content-Type, YourOwnHeader")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)
}

//-----------------------------------------------------------------------------

// This is to checking the connection is ok with the database and then connect with the database
func dbInit() {
	// Getting db config values from db.config files by processing the file through readConfig method in config_util.go
	dbConfig, err := _Util.ReadConfig("./configs/db.config")
	if err != nil {
		log.Panic(err)
	}

	// check previous commit to see the process o taking db properties from this file ( from const properties)
	dbString := "host= " + dbConfig["host"] + " user= " + dbConfig["username"] + " password= " + dbConfig["password"] + " dbname= " + dbConfig["database"] + " port= " + dbConfig["port"] + " sslmode = disable"

	db, err = gorm.Open(postgres.Open(dbString), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Connection Successful")
}

// It’s kinda important to make sure that the Entities in concern exist as tables in the connected database.
// This particular method ensures that a table named products is created on the connected database.
// func dbMigrate() {
// 	db.AutoMigrate(User{})
// 	log.Println("Database Migration Completed...")
// }

//-----------------------------------------------------------------------------

// Get All Users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

	fmt.Println(users)
}

// Create a New User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	json.NewEncoder(w).Encode(user)
}

// Check If a User Exists
func CheckUserExist(userId string) bool {
	var user User
	db.First(&user, userId)
	if user.Id == "" {
		return false
	}
	return true
}

// Get an User by an User Id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"] // "id" would be in json format that we set in User struct

	if CheckUserExist(userId) == false {
		json.NewEncoder(w).Encode("User Not Found!")
		return
	}

	var user User

	db.First(&user, userId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Update The Existing User's Data
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"] // "id" would be in json format that we set in User struct
	if CheckUserExist(userId) == false {
		json.NewEncoder(w).Encode("User Not Found")
		return
	}

	var user User
	db.First(&user, userId)
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Delete an User
func DeleteUer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := mux.Vars(r)["id"] // "id" would be in json format that we set in User struct
	if CheckUserExist(userId) == false {
		json.NewEncoder(w).Encode("User Not Found")
		return
	}

	var user User
	db.Delete(&user, userId)
	json.NewEncoder(w).Encode("Product Deleted Successfully")
}

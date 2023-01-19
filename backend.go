package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// module name is 'backend'. check go.mod for your package module name
	"backend/models"
	_UserHandler "backend/user/delivery/http"
	_Userrepo "backend/user/repository"
	_UserUsecase "backend/user/usecase"
	_Util "backend/util"

	_ "backend/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

type User models.User

var db *gorm.DB

// @title           crud application swagger
// @version         1.0
// @description     sample crud application using vuetift 3, go lang, postgres
// @termsofservice  http://swagger.io/terms/

// @contact.name   ishmam abir
// @contact.url    https://github.com/ishmamabir
// @contact.email  ishmam.cse@gmail.com

// @license.name  apache 2.0
// @license.url   http://www.apache.org/licenses/license-2.0.html

// @host      localhost:9080
// @basepath  /
func main() {
	fmt.Println("Running")
	fmt.Println("--------------------------")

	//step 1: First initialize the database connection.
	// I used gorm for database connection
	dbInit()

	router := mux.NewRouter()

	userRepo := _Userrepo.NewPsqlUserRepository(db)
	userUsecase := _UserUsecase.NewUserUsecase(userRepo)
	_UserHandler.NewUserHandler(router, userUsecase)

	// define swagger ui path and define file path
	// http://localhost:9080/swagger/index.html Swagger can be found in this link
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler).Methods(http.MethodGet)

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

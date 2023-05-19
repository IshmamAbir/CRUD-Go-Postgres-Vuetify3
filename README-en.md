## [v1: basic system of a crud operation (begineer level)](/tree/v1)

## [v2: Package System folder architecture (following Clean Architecture)](/tree/v2) 
<br>

## Startup procedure

1. Create a '```democrud```' database in postgres.

1. In Visual Studio Code, open the project. 


1. Start Command Prompt and move to the following directory of the project.  

   ```
   cd C:\['your location']\CRUD-Go-Postgres-Vuetify3

   ```

2. Start the backend  

   ```
   go run backend.go
   ```

1. Start the command prompt and move to the following directory of the project.  

   ```
   cd ./frontend
   ```

2. Launch the frontend  

   ```
   npm run serve
   ```

3. Access with a Chronium-like browser such as Edge/Safari.    
   (When npm run serve completes the build, the following URL is displayed.)  
   App running at:
   [http://localhost:8080/](http//localhost:8080/)  


## Library Used

1. Gorm ( gorm.io/gorm)
2. Gorilla Mux (github.com/gorilla/mux)
3.  Postgres (gorm.io/driver/postgres)


# Swagger Implemented here
swagger url 
```
http://localhost:9080/swagger/index.html
```

## Swagger Library
1. Swaggo ("github.com/swaggo/http-swagger")

## Running Process
1. download swag for go
```
go get github.com/swaggo/swag/cmd/swag
```

1. run swag in the root directory of the project
```
swag init 
```

1. if main method file name is not main.go, then run 
```
3. swag init -g [main method file name. e.g: backend.go]
``` 

5. download `http-swagger` using this command 
```
go get -u github.com/swaggo/http-swagger
```

1. Import the http-swagger in your main file
   ```
   import "github.com/swaggo/http-swagger" 
   ```
2. After defining any method with the swagger annotation, run the command of [2 or 3] every time.

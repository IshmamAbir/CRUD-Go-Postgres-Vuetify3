## [v1: basic system of a crud operation (begineer level)](https://github.com/IshmamAbir/CRUD-Go-Postgres-Vuetify3/tree/v1)

## [v2: Package System folder architecture (following Clean Architecture)](https://github.com/IshmamAbir/CRUD-Go-Postgres-Vuetify3/tree/v2)

<br>

![Structure](./images/folder_architecture.png)

<hr>
<br>

## Startup procedure

1. Create a '`democrud`' database in postgres.

1. In Visual Studio Code, open the project.

1. Start Command Prompt and move to the following directory of the project.

   ```
   cd C:\['your location']\CRUD-Go-Postgres-Vuetify3

   ```

1. Start the backend

   ```
   go run backend.go
   ```

1. Start the command prompt and move to the following directory of the project.

   ```
   cd ./frontend
   ```

1. Launch the frontend

   ```
   npm run serve
   ```

1. Access with a Chronium-like browser such as Edge/Safari.  
   (When npm run serve completes the build, the following URL is displayed.)  
   App running at:
   [http://localhost:8080/](http//localhost:8080/)

## Library Used

1. Gorm ( gorm.io/gorm)
2. Gorilla Mux (github.com/gorilla/mux)
3. Postgres (gorm.io/driver/postgres)

// main.go

package main

import (
	"log"
	"net/http"
	"student-api/db"
	"student-api/routes"
)

func main() {
    database := db.Migrate()

    r := routes.StudentRoutes(database)
    log.Fatal(http.ListenAndServe(":8080", r))
}

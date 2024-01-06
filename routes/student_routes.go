// routes/student_routes.go

package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"student-api/models"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

func StudentRoutes(db *gorm.DB) *chi.Mux {
    r := chi.NewRouter()
    r.Post("/students", createStudent(db))
    r.Get("/students", getStudents(db))
    return r
}

func createStudent(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var student models.Student
        if err := r.ParseForm(); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }
        student.Name = r.FormValue("name")
        student.Email = r.FormValue("email")

		body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Failed to read request body", http.StatusInternalServerError)
            return
        }
        log.Printf("Request Body: %s\n", body)

        if err := db.Create(&student).Error; err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
    }
}

func getStudents(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var students []models.Student
        if err := db.Find(&students).Error; err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        // Convert students slice to JSON
        jsonResponse, err := json.Marshal(students)
        if err != nil {
            http.Error(w, "Failed to marshal students to JSON", http.StatusInternalServerError)
            return
        }

        // Set content type and send response
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(jsonResponse)
    }
}

package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"student-api/models"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

func StudentRoutes(db *gorm.DB) *chi.Mux {
    r := chi.NewRouter()
    r.Post("/students", createStudent(db))
    r.Get("/students", getStudents(db))
	r.Put("/students/{id}", updateStudent(db))
	r.Delete("/students/{id}", deleteStudent(db))
	r.Get("/students/{id}", getStudent(db))
    return r
}

func createStudent(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        var studentReq models.StudentRequest
        if err := json.NewDecoder(r.Body).Decode(&studentReq); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }
        

        student := models.Student{Name: studentReq.Name, Email: studentReq.Email}


        if err := db.Create(&student).Error; err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
		w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(student)
    }
}

func getStudent(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Extract student ID from URL parameter
        studentIDStr := chi.URLParam(r, "id")
        studentID, err := strconv.Atoi(studentIDStr)
        if err != nil {
            http.Error(w, "Invalid student ID", http.StatusBadRequest)
            return
        }

        // Fetch specific student by their id
        var student models.Student
        if err := db.First(&student, studentID).Error; err != nil {
            http.Error(w, "Student not found", http.StatusNotFound)
            return
        }

        // Send success response with the fetched student data
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(student)
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

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(jsonResponse)
    }
}

func updateStudent(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        studentIDStr := chi.URLParam(r, "id")
        studentID, err := strconv.Atoi(studentIDStr)
        if err != nil {
            http.Error(w, "Invalid student ID", http.StatusBadRequest)
            return
        }


        var updatedStudent models.Student
        if err := json.NewDecoder(r.Body).Decode(&updatedStudent); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }


        if err := db.Model(&models.Student{}).Where("id = ?", studentID).Updates(updatedStudent).Error; err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }


        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(updatedStudent) // Send the updated student data in the response
    }
}

func deleteStudent(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Extract student ID from URL parameter
        studentIDStr := chi.URLParam(r, "id")
        studentID, err := strconv.Atoi(studentIDStr)
        if err != nil {
            http.Error(w, "Invalid student ID", http.StatusBadRequest)
            return
        }

        // Delete student record from the database
        if err := db.Delete(&models.Student{}, studentID).Error; err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }


        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Student deleted successfully"))
    }
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Student defines the structure for a single student
type Student struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Age          int      `json:"age"`
	Grade        string   `json:"grade"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	Courses      []string `json:"courses"`
	Attendance   float64  `json:"attendance"`
	EnrolledDate string   `json:"enrolledDate"`
}

// ManagementSystem holds the slice of students
type ManagementSystem struct {
	Students []Student
	NextID   int
	mu       sync.Mutex
}

var ms = ManagementSystem{
	Students: []Student{},
	NextID:   1,
}

// CORS middleware
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// Get all students
func getStudents(w http.ResponseWriter, r *http.Request) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ms.Students)
}

// Get single student by ID
func getStudent(w http.ResponseWriter, r *http.Request) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, student := range ms.Students {
		if student.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(student)
			return
		}
	}

	http.Error(w, "Student not found", http.StatusNotFound)
}

// Add new student
func addStudent(w http.ResponseWriter, r *http.Request) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	student.ID = ms.NextID
	student.EnrolledDate = time.Now().Format("2006-01-02")
	if student.Courses == nil {
		student.Courses = []string{}
	}
	if student.Attendance == 0 {
		student.Attendance = 100.0
	}

	ms.Students = append(ms.Students, student)
	ms.NextID++

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

// Delete student
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, student := range ms.Students {
		if student.ID == id {
			ms.Students = append(ms.Students[:i], ms.Students[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Student deleted successfully"})
			return
		}
	}

	http.Error(w, "Student not found", http.StatusNotFound)
}

// Add course to student
func addCourse(w http.ResponseWriter, r *http.Request) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	var data struct {
		ID     int    `json:"id"`
		Course string `json:"course"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range ms.Students {
		if ms.Students[i].ID == data.ID {
			ms.Students[i].Courses = append(ms.Students[i].Courses, data.Course)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ms.Students[i])
			return
		}
	}

	http.Error(w, "Student not found", http.StatusNotFound)
}

// Update attendance
func updateAttendance(w http.ResponseWriter, r *http.Request) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	var data struct {
		ID         int     `json:"id"`
		Attendance float64 `json:"attendance"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range ms.Students {
		if ms.Students[i].ID == data.ID {
			ms.Students[i].Attendance = data.Attendance
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ms.Students[i])
			return
		}
	}

	http.Error(w, "Student not found", http.StatusNotFound)
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// API routes
	http.HandleFunc("/api/students", enableCORS(getStudents))
	http.HandleFunc("/api/student", enableCORS(getStudent))
	http.HandleFunc("/api/add", enableCORS(addStudent))
	http.HandleFunc("/api/delete", enableCORS(deleteStudent))
	http.HandleFunc("/api/addCourse", enableCORS(addCourse))
	http.HandleFunc("/api/updateAttendance", enableCORS(updateAttendance))

	fmt.Println("🚀 Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

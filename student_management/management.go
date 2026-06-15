package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Student defines the structure for a single student
type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

// ManagementSystem holds the slice of students
type ManagementSystem struct {
	Students []Student
	NextID   int
}

// AddStudent creates a new student and adds it to the system
func (ms *ManagementSystem) AddStudent(name string, age int, grade string) {
	newStudent := Student{
		ID:    ms.NextID,
		Name:  name,
		Age:   age,
		Grade: grade,
	}
	ms.Students = append(ms.Students, newStudent)
	fmt.Printf("🎉 Student '%s' added successfully with ID: %d\n", name, ms.NextID)
	ms.NextID++ // Increment ID for the next student
}

// ViewStudents displays all students in a formatted way
func (ms *ManagementSystem) ViewStudents() {
	if len(ms.Students) == 0 {
		fmt.Println("⚠️ No students found in the system.")
		return
	}

	fmt.Println("\n--- Student List ---")
	fmt.Printf("%-5s %-20s %-5s %-5s\n", "ID", "Name", "Age", "Grade")
	fmt.Println(strings.Repeat("-", 40))
	for _, s := range ms.Students {
		fmt.Printf("%-5d %-20s %-5d %-5s\n", s.ID, s.Name, s.Age, s.Grade)
	}
}

// DeleteStudent removes a student by their ID
func (ms *ManagementSystem) DeleteStudent(id int) {
	for i, s := range ms.Students {
		if s.ID == id {
			// Remove from slice by appending elements before and after the index
			ms.Students = append(ms.Students[:i], ms.Students[i+1:]...)
			fmt.Printf("🗑️ Student with ID %d deleted successfully.\n", id)
			return
		}
	}
	fmt.Printf("❌ Student with ID %d not found.\n", id)
}

func main() {
	// Initialize the system
	ms := ManagementSystem{
		Students: []Student{},
		NextID:   1,
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Student Management System ===")
		fmt.Println("1. Add Student")
		fmt.Println("2. View Students")
		fmt.Println("3. Delete Student")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Enter Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter Age: ")
			ageStr, _ := reader.ReadString('\n')
			age, _ := strconv.Atoi(strings.TrimSpace(ageStr))

			fmt.Print("Enter Grade (e.g., A, B, C): ")
			grade, _ := reader.ReadString('\n')
			grade = strings.TrimSpace(grade)

			ms.AddStudent(name, age, grade)

		case "2":
			ms.ViewStudents()

		case "3":
			fmt.Print("Enter Student ID to delete: ")
			idStr, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idStr))
			ms.DeleteStudent(id)

		case "4":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("❌ Invalid option, please try again.")
		}
	}
}
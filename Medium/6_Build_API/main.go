package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//DB

var courses []Course

// middleware

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - Practice")
	router := mux.NewRouter()
	courses = append(courses, Course{CourseId: "1", CourseName: "React JS", CoursePrice: 299, Author: &Author{Fullname: "Dejbit Pramanick", Website: "dev.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Node JS", CoursePrice: 499, Author: &Author{Fullname: "Dejbit Pramanick", Website: "dev.com"}})

	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses/get", getAllCourses).Methods("GET")
	router.HandleFunc("/courses/{id}", getCourseByID).Methods("GET")
	router.HandleFunc("/courses/create", createCourse).Methods("POST")
	router.HandleFunc("/courses/update/{id}", udpateCourse).Methods("PUT")
	router.HandleFunc("/courses/delete/{id}", deleteCourse).Methods("DELETE")

	// Listen a port
	log.Fatal(http.ListenAndServe(":4000", router))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<p>Server is running.</p>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[ Get All Courses ]")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourseByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[ Get Courses By ID ]")
	w.Header().Set("Content-Type", "application/json")

	// Grab ID from request
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given ID.")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[ Create new course ]")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data.")
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON.")
		return
	}

	// Create ID 

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func udpateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[ Update course ]")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given ID.")
	return
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[ Update course ]")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("No course found with given ID.")
	return
}
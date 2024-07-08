package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Students []Student `json:"students"`
}
type Student struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering into students handler")
	var response Response
	students := prepareResponse()
	response.Students = students

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering HealthCheckHandler to check health of end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func prepareResponse() []Student {
	var students []Student
	var student Student

	student = Student{1, "Prakhar", "Sharma"}
	students = append(students, student)

	student = Student{
		Id:        2,
		FirstName: "Piyush",
		LastName:  "Bajaj",
	}
	students = append(students, student)

	student.Id = 3
	student.FirstName = "Viraj"
	student.LastName = "Chavan"
	students = append(students, student)

	return students

}

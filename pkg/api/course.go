package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/tonymontanapaffpaff/golang-training-university/pkg/data"

	"github.com/gorilla/mux"
)

type courseAPI struct {
	data *data.CourseData
}

func ServeCourseResource(r *mux.Router, data data.CourseData) {
	api := &courseAPI{data: &data}
	r.HandleFunc("/courses", api.getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{code}", api.getCourse).Methods("GET")
	r.HandleFunc("/courses", api.createCourse).Methods("POST")
	r.HandleFunc("/courses/{code}", api.updateCourseDescription).Methods("PATCH")
	r.HandleFunc("/courses/{code}", api.deleteCourse).Methods("DELETE")
}

func (a courseAPI) getAllCourses(writer http.ResponseWriter, request *http.Request) {
	courses, err := a.data.ReadAll()
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get list of courses"))
		if err != nil {
			log.Println(err)
		}
		return
	}
	err = json.NewEncoder(writer).Encode(courses)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a courseAPI) getCourse(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	code, err := strconv.Atoi(vars["code"])
	if err != nil {
		_, err := writer.Write([]byte("course code must be a number"))
		if err != nil {
			log.Println(err)
		}
		return
	}
	course, err := a.data.Read(code)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get course"))
		if err != nil {
			log.Println(err)
		}
		return
	}
	err = json.NewEncoder(writer).Encode(course)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a courseAPI) createCourse(writer http.ResponseWriter, request *http.Request) {
	course := new(data.Course)
	err := json.NewDecoder(request.Body).Decode(&course)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if course == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = a.data.Add(*course)
	if err != nil {
		log.Println("course hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (a courseAPI) updateCourseDescription(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	code, err := strconv.Atoi(vars["code"])
	if err != nil {
		_, err := writer.Write([]byte("course code must be a number"))
		if err != nil {
			log.Println(err)
		}
		return
	}
	var description string
	err = json.NewDecoder(request.Body).Decode(&description)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if description == "" {
		log.Printf("empty data\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = a.data.ChangeDescription(code, description)
	if err != nil {
		log.Println("course hasn't been updated")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (a courseAPI) deleteCourse(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	code, err := strconv.Atoi(vars["code"])
	if err != nil {
		_, err := writer.Write([]byte("course code must be a number"))
		if err != nil {
			log.Println(err)
		}
		return
	}
	err = a.data.Delete(code)
	if err != nil {
		log.Println("course hasn't been deleted")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

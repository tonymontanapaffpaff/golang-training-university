package api

import (
	"encoding/json"
	"net/http"

	"github.com/tonymontanapaffpaff/golang-training-university/pkg/data"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type courseAPI struct {
	data *data.CourseData
}

func ServeCourseResource(r *mux.Router, data data.CourseData) {
	api := &courseAPI{data: &data}
	r.HandleFunc("/courses", api.getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", api.getCourse).Methods("GET")
	r.HandleFunc("/courses", api.createCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", api.updateCourseDescription).Methods("PATCH")
	r.HandleFunc("/courses/{id}", api.deleteCourse).Methods("DELETE")
}

func (a courseAPI) getAllCourses(writer http.ResponseWriter, request *http.Request) {
	courses, err := a.data.ReadAll()
	if err != nil {
		log.Error(err)
		_, err := writer.Write([]byte("got an error when tried to get list of courses"))
		if err != nil {
			log.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	log.Info("GetAllCourses method result:")
	for _, value := range courses {
		log.Info(value)
	}
	err = json.NewEncoder(writer).Encode(courses)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a courseAPI) getCourse(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	if len(id) > 24 {
		log.Error("Error: can't get course, course id must be less then 25 symbols")
		_, err := writer.Write([]byte("course id must be less then 25 symbols"))
		if err != nil {
			log.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	course, err := a.data.Read(id)
	if err != nil {
		log.Error(err)
		_, err := writer.Write([]byte("got an error when tried to get course"))
		if err != nil {
			log.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	log.Infof("GetCourse method result: %v", course)
	err = json.NewEncoder(writer).Encode(course)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a courseAPI) createCourse(writer http.ResponseWriter, request *http.Request) {
	course := new(data.Course)
	err := json.NewDecoder(request.Body).Decode(&course)
	if err != nil {
		log.Errorf("failed reading JSON: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if course == nil {
		log.Error("empty JSON")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := a.data.Add(*course)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Infof("CreateCourse method result: %s", result)
	writer.WriteHeader(http.StatusCreated)
}

type Description struct {
	Description string `json:"description"`
}

func (a courseAPI) updateCourseDescription(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	if len(id) > 24 {
		log.Error("Error: can't update course description, course id must be less then 25 symbols")
		_, err := writer.Write([]byte("course id must be less then 25 symbols"))
		if err != nil {
			log.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	var description Description
	err := json.NewDecoder(request.Body).Decode(&description)
	if err != nil {
		log.Errorf("failed reading JSON: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if description.Description == "" {
		log.Error("empty data")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := a.data.ChangeDescription(id, description.Description)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Infof("UpdateCourseDescription method result: %s", result)
	writer.WriteHeader(http.StatusCreated)
}

func (a courseAPI) deleteCourse(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	if len(id) > 24 {
		log.Error("Error: can't delete course, course id must be less then 25 symbols")
		_, err := writer.Write([]byte("course id must be not greater then 25 symbols"))
		if err != nil {
			log.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	err := a.data.Delete(id)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Info("DeleteCourse method successfully done")
	writer.WriteHeader(http.StatusCreated)
}

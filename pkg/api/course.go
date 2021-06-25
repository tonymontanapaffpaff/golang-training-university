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

func NewCourseAPI(data *data.CourseData) courseAPI {
	return courseAPI{data: data}
}

func ServeCourseResource(r *mux.Router, data data.CourseData) {
	api := &courseAPI{data: &data}
	r.HandleFunc("/courses", api.GetAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", api.GetCourse).Methods("GET")
	r.HandleFunc("/courses", api.CreateCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", api.UpdateCourseDescription).Methods("PATCH")
	r.HandleFunc("/courses/{id}", api.DeleteCourse).Methods("DELETE")
}

func (a courseAPI) GetAllCourses(writer http.ResponseWriter, request *http.Request) {
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
	log.Debugf("GetAllCourses method result: %v", courses)
	err = json.NewEncoder(writer).Encode(courses)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a courseAPI) GetCourse(writer http.ResponseWriter, request *http.Request) {
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
	log.Debugf("GetCourse method result: %v", course)
	err = json.NewEncoder(writer).Encode(course)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a courseAPI) CreateCourse(writer http.ResponseWriter, request *http.Request) {
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
	log.Debugf("CreateCourse method result: %s", result)
	writer.WriteHeader(http.StatusCreated)
}

type Description struct {
	Description string `json:"description"`
}

func (a courseAPI) UpdateCourseDescription(writer http.ResponseWriter, request *http.Request) {
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
	log.Debugf("UpdateCourseDescription method result: %s", result)
	writer.WriteHeader(http.StatusCreated)
}

func (a courseAPI) DeleteCourse(writer http.ResponseWriter, request *http.Request) {
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
	log.Debug("DeleteCourse method successfully done")
	writer.WriteHeader(http.StatusCreated)
}

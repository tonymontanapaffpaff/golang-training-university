package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/tonymontanapaffpaff/golang-training-university/pkg/data"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type StudentAPI struct {
	data *data.StudentData
}

func ServeStudentResource(r *mux.Router, data data.StudentData) {
	api := &StudentAPI{data: &data}
	r.HandleFunc("/students", api.GetAllStudents).Methods("GET")
	r.HandleFunc("/students/{id}", api.GetStudent).Methods("GET")
	r.HandleFunc("/students", api.CreateStudent).Methods("POST")
	r.HandleFunc("/students/{id}", api.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/students/{id}", api.MakePayment).Methods("POST")
}

func (a StudentAPI) GetAllStudents(writer http.ResponseWriter, request *http.Request) {
	students, err := a.data.ReadAll()
	if err != nil {
		log.Error(err)
		_, err = writer.Write([]byte("got an error when tried to get list of students"))
		if err != nil {
			log.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	log.Debugf("GetAllStudents method result: %v", students)
	err = json.NewEncoder(writer).Encode(students)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a StudentAPI) GetStudent(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		_, err = writer.Write([]byte("student id must be a number"))
		if err != nil {
			log.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	student, err := a.data.Read(id)
	if err != nil {
		log.Error(err)
		_, err = writer.Write([]byte("got an error when tried to get student"))
		if err != nil {
			log.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	log.Debugf("GetStudent method result: %v", student)
	err = json.NewEncoder(writer).Encode(student)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a StudentAPI) CreateStudent(writer http.ResponseWriter, request *http.Request) {
	student := new(data.Student)
	err := json.NewDecoder(request.Body).Decode(&student)
	if err != nil {
		log.Errorf("failed reading JSON: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if student == nil {
		log.Error("empty JSON")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := a.data.Add(*student)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Debugf("CreateStudent method result: %d", result)
	writer.WriteHeader(http.StatusCreated)
}

func (a StudentAPI) DeleteStudent(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		_, err = writer.Write([]byte("student id must be a number"))
		if err != nil {
			log.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = a.data.Delete(id)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Debug("DeleteStudent method successfully done")
	writer.WriteHeader(http.StatusCreated)
}

type PaymentRequisites struct {
	CourseId int `json:"course_id"`
	Payment  int `json:"payment"`
}

func (a StudentAPI) MakePayment(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	studentId, err := strconv.Atoi(vars["id"])
	if err != nil {
		_, err = writer.Write([]byte("studentId must be a number"))
		if err != nil {
			log.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	paymentRequisites := new(PaymentRequisites)
	err = json.NewDecoder(request.Body).Decode(&paymentRequisites)
	if err != nil {
		log.Errorf("failed reading JSON: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if paymentRequisites == nil {
		log.Error("empty JSON")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = a.data.Pay(studentId, paymentRequisites.CourseId, paymentRequisites.Payment)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}

	log.Debug("DeleteStudent method successfully done")
	writer.WriteHeader(http.StatusCreated)
	_, err = writer.Write([]byte("payment successful"))
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

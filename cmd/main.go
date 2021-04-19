package main

import (
	"log"
	"os"

	"github.com/golang-training-university/pkg/data"
	"github.com/golang-training-university/pkg/db"
)

var (
	host     = os.Getenv("DB_USERS_HOST")
	port     = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_DBNAME")
	password = os.Getenv("DB_USERS_PASSWORD")
	sslmode  = os.Getenv("DB_USERS_SSL")
)

func init() {
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "university"
	}
	if password == "" {
		password = "postgres"
	}
	if sslmode == "" {
		sslmode = "disable"
	}
}

func main() {
	conn, err := db.GetConnection(host, port, user, dbname, password, sslmode)
	if err != nil {
		log.Fatalf("can't connect to database, error: %v", err)
	}

	studentData := data.NewStudentData(conn)

	//newStudent := data.Student{
	//	Id:        20211101,
	//	FirstName: "Kyle",
	//	LastName:  "Kuzma",
	//	IsActive:  false,
	//}

	//id, err := studentData.Add(newStudent)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println("Inserted newStudent id is:", id)
	//
	//err = studentData.Delete(20211101)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println("Successfully deletion")

	//students, err := studentData.ReadAll()
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(students)
	//
	//student, err := studentData.Read(20174201)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println("Searching result:", student)

	changedUserId, err := studentData.ChangeStatus(20174201)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Student status with id=%d is changed\n", changedUserId)

	//currentRate, err := studentData.GetCurrentRate(student.Id)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Printf("Current rate for student with id=%d: %v\n", student.Id, currentRate)
	//
	//coursesList, err := studentData.GetCoursesList(student.Id)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Printf("List of courses for student with id=%d: %v\n", student.Id, coursesList)
	//
	//courseData := data.NewCourseData(conn)
	//courses, err := courseData.ReadAll()
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Printf("List of courses: %v\n", courses)
	//
	//departmentName, err := courseData.GetDepartmentName(202)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Printf("Department name for course with code=%d: %v\n", 202, departmentName)
}

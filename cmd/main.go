package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/tonymontanapaffpaff/golang-training-university/pkg/api"
	"github.com/tonymontanapaffpaff/golang-training-university/pkg/data"
	"github.com/tonymontanapaffpaff/golang-training-university/pkg/db"

	"github.com/gorilla/mux"
)

var (
	serverEndpoint = os.Getenv("SERVER_ENDPOINT")
	dbHost         = os.Getenv("DB_HOST")
	dbPort         = os.Getenv("DB_PORT")
	dbUser         = os.Getenv("DB_USER")
	dbName         = os.Getenv("DB_NAME")
	dbPassword     = os.Getenv("DB_PASSWORD")
	sslMode        = os.Getenv("SSL_MODE")
)

func init() {
	if serverEndpoint == "" {
		serverEndpoint = "8080"
	}
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "5432"
	}
	if dbUser == "" {
		dbUser = "postgres"
	}
	if dbName == "" {
		dbName = "university"
	}
	if dbPassword == "" {
		dbPassword = "root"
	}
	if sslMode == "" {
		sslMode = "disable"
	}
}

func main() {
	conn, err := db.GetConnection(dbHost, dbPort, dbUser, dbName, dbPassword, sslMode)
	if err != nil {
		log.Fatalf("can't connect to database, error: %v", err)
	}

	r := mux.NewRouter()
	courseData := data.NewCourseData(conn)
	api.ServeCourseResource(r, *courseData)
	studentData := data.NewStudentData(conn)
	api.ServeStudentResource(r, *studentData)
	paymentData := data.NewPaymentData(conn)
	api.ServePaymentResource(r, *paymentData)
	r.Use(mux.CORSMethodMiddleware(r))

	listener, err := net.Listen("tcp", ":"+serverEndpoint)
	if err != nil {
		log.Fatal("Server Listen port...")
	}

	err = http.Serve(listener, r)
	if err != nil {
		log.Fatal("Server has been crashed...")
	}
}

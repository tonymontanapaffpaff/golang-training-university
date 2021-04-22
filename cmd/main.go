package main

import (
	"log"
	"net"
	"net/http"

	"github.com/tonymontanapaffpaff/golang-training-university/config"
	"github.com/tonymontanapaffpaff/golang-training-university/pkg/api"
	"github.com/tonymontanapaffpaff/golang-training-university/pkg/data"
	"github.com/tonymontanapaffpaff/golang-training-university/pkg/db"

	"github.com/gorilla/mux"
)

func main() {
	appConfig := config.GetConfig()
	conn, err := db.GetConnection(appConfig)
	if err != nil {
		log.Fatalf("can't connect to database, error: %v", err)
	}

	r := mux.NewRouter()
	courseData := data.NewCourseData(conn)
	api.ServeCourseResource(r, *courseData)
	r.Use(mux.CORSMethodMiddleware(r))

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Server Listen port...")
	}

	err = http.Serve(listener, r)
	if err != nil {
		log.Fatal("Server has been crashed...")
	}
}

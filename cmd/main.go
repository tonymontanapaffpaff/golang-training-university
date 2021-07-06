package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/tonymontanapaffpaff/golang-training-university/pkg/api"
	"github.com/tonymontanapaffpaff/golang-training-university/pkg/data"
	"github.com/tonymontanapaffpaff/golang-training-university/pkg/db"

	"github.com/gorilla/mux"
)

var (
	serverEndpoint = os.Getenv("SERVER_ENDPOINT")
	dbHost         = os.Getenv("DB_USERS_HOST")
	dbPort         = os.Getenv("DB_USERS_PORT")
)

func init() {
	if serverEndpoint == "" {
		serverEndpoint = ":8080"
	}
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "27017"
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := db.GetClient(ctx, dbHost, dbPort)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	r := mux.NewRouter()
	courseData := data.NewCourseData(client.Database("university").Collection("courses"))
	api.ServeCourseResource(r, *courseData)
	r.Use(mux.CORSMethodMiddleware(r))

	listener, err := net.Listen("tcp", serverEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	err = http.Serve(listener, r)
	if err != nil {
		log.Fatal(err)
	}
}

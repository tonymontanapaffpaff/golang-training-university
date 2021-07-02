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
	"github.com/tonymontanapaffpaff/golang-training-university/pkg/data/auth"
	"github.com/tonymontanapaffpaff/golang-training-university/pkg/db"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var (
	dbHost         = os.Getenv("MONGO_HOST")
	dbPort         = os.Getenv("MONGO_PORT")
	redisHost      = os.Getenv("REDIS_HOST")
	redisPort      = os.Getenv("REDIS_PORT")
	redisPassword  = os.Getenv("REDIS_PASSWORD")
	serverEndpoint = os.Getenv("SERVER_ENDPOINT")
)

func init() {
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "27017"
	}
	if redisHost == "" {
		redisHost = "localhost"
	}
	if redisPort == "" {
		redisPort = "6379"
	}
	if serverEndpoint == "" {
		serverEndpoint = ":8080"
	}
}

func NewRedisDB(host, port, password string) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	return redisClient
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := db.GetClient(ctx, dbHost, dbPort)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	redisClient := NewRedisDB(redisHost, redisPort, redisPassword)
	IAuth := auth.NewAuth(redisClient)
	IToken := auth.NewToken()

	r := mux.NewRouter()
	courseData := data.NewCourseData(client.Database("university").Collection("courses"), IAuth, IToken)
	userData := data.NewUserData(client.Database("university").Collection("users"), IAuth, IToken)
	api.ServeCourseResource(r, *courseData)
	api.ServeUserResource(r, *userData)
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

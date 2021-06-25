package test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/tonymontanapaffpaff/golang-training-university/pkg/data"
	"github.com/tonymontanapaffpaff/golang-training-university/pkg/db"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var courses = map[string]data.Course{
	"First course": {
		ID:           primitive.NewObjectID(),
		Title:        "First course title",
		DepartmentId: primitive.ObjectID{},
		Description:  "",
	},
	"Second course": {
		ID:           primitive.NewObjectID(),
		Title:        "Second course title",
		DepartmentId: primitive.ObjectID{},
		Description:  "",
	},
	"Third course": {
		ID:           primitive.NewObjectID(),
		Title:        "Third course title",
		DepartmentId: primitive.ObjectID{},
		Description:  "",
	},
}

func insertData(ctx context.Context, collection *mongo.Collection) {
	for _, course := range courses {
		if _, err := collection.InsertOne(ctx, course); err != nil {
			log.Error(err)
		}
	}
}

func TestMain(m *testing.M) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := db.GetClient(ctx, "localhost", "27017")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("university").Collection("courses")

	collection.Drop(ctx)
	insertData(ctx, collection)

	// Run the test suite
	retCode := m.Run()

	os.Exit(retCode)
}

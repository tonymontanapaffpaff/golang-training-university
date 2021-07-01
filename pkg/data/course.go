package data

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/tonymontanapaffpaff/golang-training-university/pkg/data/auth"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Course struct {
	ID           primitive.ObjectID `json:"ID" bson:"_id"`
	Title        string             `json:"title" bson:"title"`
	DepartmentId primitive.ObjectID `json:"department_id" bson:"department_id"`
	Description  string             `json:"description" bson:"description"`
}

type CourseData struct {
	Collection *mongo.Collection
	IAuth      auth.IAuth
	IToken     auth.IToken
}

func NewCourseData(collection *mongo.Collection, IAuth auth.IAuth, IToken auth.IToken) *CourseData {
	return &CourseData{Collection: collection, IAuth: IAuth, IToken: IToken}
}

func (d CourseData) Add(course Course, request *http.Request) (string, HttpErr) {
	metadata, err := d.IToken.ExtractTokenMetadata(request)
	if err != nil {
		return "", HttpErr{
			Err:        fmt.Errorf("unauthorized"),
			StatusCode: http.StatusUnauthorized,
		}
	}
	_, err = d.IAuth.FetchAuth(metadata.TokenUuid)
	if err != nil {
		return "", HttpErr{
			Err:        fmt.Errorf("unauthorized"),
			StatusCode: http.StatusUnauthorized,
		}
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = d.Collection.InsertOne(ctx, course)
	if err != nil {
		return "-1", HttpErr{
			Err:        fmt.Errorf("can't create course, error: %w", err),
			StatusCode: http.StatusBadRequest,
		}
	}
	return course.ID.String(), HttpErr{}
}

func (d CourseData) Read(id string) (Course, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.D{{"_id", objectID}}
	var result Course

	err = d.Collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, fmt.Errorf("can't read course with given ID, error: %w", err)
	}
	return result, nil
}

func (d CourseData) ReadAll() ([]*Course, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{}
	opt := options.Find()
	var results []*Course

	result, err := d.Collection.Find(ctx, filter, opt)
	if err != nil {
		return nil, fmt.Errorf("can't read courses from database, error: %w", err)
	}
	defer result.Close(ctx)

	for result.Next(ctx) {
		var elem Course
		err := result.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := result.Err(); err != nil {
		log.Fatal(err)
	}

	return results, nil
}

func (d CourseData) ChangeDescription(id string, description string, request *http.Request) (string, HttpErr) {
	metadata, err := d.IToken.ExtractTokenMetadata(request)
	if err != nil {
		return "", HttpErr{
			Err:        fmt.Errorf("unauthorized"),
			StatusCode: http.StatusUnauthorized,
		}
	}
	_, err = d.IAuth.FetchAuth(metadata.TokenUuid)
	if err != nil {
		return "", HttpErr{
			Err:        fmt.Errorf("unauthorized"),
			StatusCode: http.StatusUnauthorized,
		}
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.D{primitive.E{
		Key:   "_id",
		Value: objectID,
	}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{
			Key:   "description",
			Value: description,
		},
	}}}

	_, err = d.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return "-1", HttpErr{
			Err:        fmt.Errorf("can't update course description, error: %w", err),
			StatusCode: http.StatusBadRequest,
		}
	}
	return fmt.Sprintf("ID: %s, description: %s", id, description), HttpErr{}
}

func (d CourseData) Delete(id string, request *http.Request) HttpErr {
	metadata, err := d.IToken.ExtractTokenMetadata(request)
	if err != nil {
		return HttpErr{
			Err:        fmt.Errorf("unauthorized"),
			StatusCode: http.StatusUnauthorized,
		}
	}
	_, err = d.IAuth.FetchAuth(metadata.TokenUuid)
	if err != nil {
		return HttpErr{
			Err:        fmt.Errorf("unauthorized"),
			StatusCode: http.StatusUnauthorized,
		}
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.D{{"_id", objectID}}

	_, err = d.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return HttpErr{
			Err:        fmt.Errorf("can't delete course from database, error: %w", err),
			StatusCode: http.StatusBadRequest,
		}
	}
	return HttpErr{}
}

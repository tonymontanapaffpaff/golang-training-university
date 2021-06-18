package data

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Course struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Title        string             `json:"title" bson:"title"`
	DepartmentId primitive.ObjectID `json:"department_id" bson:"department_id"`
	Description  string             `json:"description" bson:"description"`
}

type CourseData struct {
	Collection *mongo.Collection
}

func NewCourseData(collection *mongo.Collection) *CourseData {
	return &CourseData{Collection: collection}
}

func (u CourseData) Add(course Course) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := u.Collection.InsertOne(ctx, course)
	if err != nil {
		return "-1", fmt.Errorf("can't create course, error: %w", err)
	}
	return course.ID.String(), nil
}

func (u CourseData) Read(id string) (Course, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.D{{"_id", id}}
	var result Course

	err := u.Collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, fmt.Errorf("can't read course with given id, error: %w", err)
	}
	return result, nil
}

func (u CourseData) ReadAll() ([]*Course, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{}
	opt := options.Find()
	var results []*Course

	result, err := u.Collection.Find(ctx, filter, opt)
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

func (u CourseData) ChangeDescription(id string, description string) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.D{primitive.E{
		Key:   "_id",
		Value: id,
	}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{
			Key:   "description",
			Value: description,
		},
	}}}

	_, err := u.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return "-1", fmt.Errorf("can't update course description, error: %w", err)
	}
	return fmt.Sprintf("id: %s, description: %s", id, description), nil
}

func (u CourseData) Delete(id string) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.D{{"_id", id}}

	_, err := u.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("can't delete course from database, error: %w", err)
	}
	return nil
}

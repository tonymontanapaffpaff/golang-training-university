// +build integration

package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"github.com/tonymontanapaffpaff/golang-training-university/pkg/data"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestHandleGetAllPages(t *testing.T) {
	// Create the requestURI
	requestURL := url.URL{
		Scheme:      "http",
		Opaque:      "",
		User:        nil,
		Host:        "localhost:8080",
		Path:        "courses",
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    "",
		Fragment:    "",
		RawFragment: "",
	}

	var client http.Client
	resp, err := client.Get(requestURL.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Status should be http.StatusOK
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var testResults []map[string]interface{}

	// The response body should be unmarshalled
	errJSON := json.NewDecoder(resp.Body).Decode(&testResults)
	if errJSON != nil {
		t.Errorf("Error unmarshalling JSON %s",
			errJSON.Error())
	}

	// The amount of returned courses should be the same as the amount of created
	if len(testResults) != len(courses) {
		t.Errorf(`Expected "%d" results, got "%d"`, len(courses), len(testResults))
	}
}

func TestHandleGetPage(t *testing.T) {
	course := courses["First course"]

	requestURL := url.URL{
		Scheme:      "http",
		Opaque:      "",
		User:        nil,
		Host:        "localhost:8080",
		Path:        "courses/" + course.ID.Hex(),
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    "",
		Fragment:    "",
		RawFragment: "",
	}

	var client http.Client
	resp, err := client.Get(requestURL.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Status should be http.StatusOK
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	testResults := make(map[string]interface{})

	// The response body should be unmarshalled
	errJSON := json.NewDecoder(resp.Body).Decode(&testResults)
	if errJSON != nil {
		t.Errorf("Error unmarshalling JSON %s",
			errJSON.Error())
	}

	// The element id should be set
	if _, ok := testResults["id"]; !ok {
		t.Fatalf(`Expected element "id", didn't get it: "%s"`, testResults["id"])
	}

	// The element title should be the same as first course title
	if testResults["title"] != course.Title {
		t.Errorf(`Expected name "%s", got "%s"`, course.Title, testResults["title"])
	}
}

func TestHandlePostPage(t *testing.T) {
	course := data.Course{
		ID:           primitive.NewObjectID(),
		Title:        "",
		DepartmentId: primitive.ObjectID{},
		Description:  "",
	}

	requestURL := url.URL{
		Scheme:      "http",
		Opaque:      "",
		User:        nil,
		Host:        "localhost:8080",
		Path:        "courses",
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    "",
		Fragment:    "",
		RawFragment: "",
	}

	b := new(bytes.Buffer)

	// The response body should be unmarshalled
	errJSON := json.NewEncoder(b).Encode(&course)
	if errJSON != nil {
		t.Errorf("Error unmarshalling JSON %s",
			errJSON.Error())
	}

	var client http.Client
	resp, err := client.Post(requestURL.String(), "application/json; charset=utf-8", b)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Status should be http.StatusCreated
	if status := resp.StatusCode; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHandlePatchPage(t *testing.T) {
	course := courses["First course"]

	requestURL := url.URL{
		Scheme:      "http",
		Opaque:      "",
		User:        nil,
		Host:        "localhost:8080",
		Path:        "courses/" + course.ID.Hex(),
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    "",
		Fragment:    "",
		RawFragment: "",
	}

	description, err := json.Marshal(map[string]interface{}{
		"description": "new description",
	})
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPatch, requestURL.String(), bytes.NewBuffer(description))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Status should be http.StatusCreated
	if status := resp.StatusCode; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHandleDeletePage(t *testing.T) {
	course := courses["First course"]

	requestURL := url.URL{
		Scheme:      "http",
		Opaque:      "",
		User:        nil,
		Host:        "localhost:8080",
		Path:        "courses/" + course.ID.Hex(),
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    "",
		Fragment:    "",
		RawFragment: "",
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, requestURL.String(), nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// Status should be http.StatusCreated
	if status := resp.StatusCode; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHandleGetPage404(t *testing.T) {
	requestURL := url.URL{
		Scheme:      "http",
		Opaque:      "",
		User:        nil,
		Host:        "localhost:8080",
		Path:        "students",
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    "",
		Fragment:    "",
		RawFragment: "",
	}

	var client http.Client
	resp, err := client.Get(requestURL.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Status should be http.StatusNotFound
	if status := resp.StatusCode; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

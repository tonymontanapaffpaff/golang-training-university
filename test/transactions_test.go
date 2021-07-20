// +build integration

package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/tonymontanapaffpaff/golang-training-university/pkg/data"
)

var studentsId = []string{"20174201", "20174202", "20174203", "20174204", "20174205"}
var coursesId = []int{207, 208, 209, 202, 203}

func TestOneFailed(t *testing.T) {
	// Create the requestURI
	requestURL := url.URL{
		Scheme:      "http",
		Opaque:      "",
		User:        nil,
		Host:        "localhost:8080",
		Path:        "students/" + studentsId[0],
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    "",
		Fragment:    "",
		RawFragment: "",
	}

	postBody, _ := json.Marshal(map[string]interface{}{
		"course_id": coursesId[0],
		"payment":   30,
	})

	responseBody := bytes.NewBuffer(postBody)

	var client http.Client
	resp, err := client.Post(requestURL.String(), "application/json; charset=UTF-8", responseBody)
	if err != nil {
		t.Fatal(err)
	}

	// Status should be http.StatusOK
	if status := resp.StatusCode; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Can't parse response body, %v", err)
	}

	expected := "insufficient funds"
	founded := string(bodyBytes)
	if founded != expected {
		t.Errorf(`Expected "%s" results, got "%s"`, expected, founded)
	}

	resp.Body.Close()

	requestURL.Path = "payments"
	resp, err = client.Get(requestURL.String())
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var payments []data.Payment
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Can't parse response body, %v", err)
	}

	err = json.Unmarshal(bodyBytes, &payments)
	if err != nil {
		t.Errorf("Can't unmarshal response body, %v", err)
	}

	foundedStatus := payments[0].Passed
	if foundedStatus != false {
		t.Errorf(`Expected "%t" results, got "%t"`, false, foundedStatus)
	}
}

func TestAllSucceed(t *testing.T) {
	// Create the requestURI
	requestURL := url.URL{
		Scheme:      "http",
		Opaque:      "",
		User:        nil,
		Host:        "localhost:8080",
		Path:        "",
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    "",
		Fragment:    "",
		RawFragment: "",
	}

	postBody, _ := json.Marshal(map[string]interface{}{
		"course_id": coursesId[0],
		"payment":   150,
	})

	var client http.Client
	responseBody := bytes.NewBuffer(postBody)
	passed := true

	for i := 0; i < 5; i++ {
		go func(i int) {
			id := "2017420" + strconv.Itoa(i+1)
			requestURL.Path = "students/" + id
			resp, err := client.Post(requestURL.String(), "application/json; charset=UTF-8", responseBody)
			if err != nil {
				t.Errorf("can't do post request, err: %v", err)
			}
			defer resp.Body.Close()

			// Status should be http.StatusOK
			if status := resp.StatusCode; status != http.StatusOK {
				passed = false
			}
		}(i)
	}

	if passed != true {
		t.Error("not all requests have been post")
	}
}

func TestRollback(t *testing.T) {
	// Create the requestURI
	requestURL := url.URL{
		Scheme:      "http",
		Opaque:      "",
		User:        nil,
		Host:        "localhost:8080",
		Path:        "",
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    "",
		Fragment:    "",
		RawFragment: "",
	}

	postBody, _ := json.Marshal(map[string]interface{}{
		"course_id": coursesId[0],
		"payment":   150,
	})

	var client http.Client
	passed := true

	for i := 0; i < 5; i++ {
		go func(i int) {
			if i == 3 {
				postBody, _ = json.Marshal(map[string]interface{}{
					"course_id": coursesId[1],
					"payment":   10,
				})
			}
			responseBody := bytes.NewBuffer(postBody)
			id := "2017420" + strconv.Itoa(i+1)
			requestURL.Path = "students/" + id
			resp, err := client.Post(requestURL.String(), "application/json; charset=UTF-8", responseBody)
			if err != nil {
				t.Errorf("can't do post request, err: %v", err)
			}
			defer resp.Body.Close()

			// Status should be http.StatusOK
			if status := resp.StatusCode; status != http.StatusOK {
				passed = false
			}
		}(i)
	}

	if passed != true {
		t.Error("not all requests worked correctly")
	}

	requestURL.Path = "students/20174204"
	resp, err := client.Get(requestURL.String())
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var student data.Student
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Can't parse response body, %v", err)
	}

	err = json.Unmarshal(bodyBytes, &student)
	if err != nil {
		t.Errorf("Can't unmarshal response body, %v", err)
	}

	foundedStatus := student.IsActive
	if foundedStatus != false {
		t.Errorf(`Expected "%t" results, got "%t"`, false, foundedStatus)
	}
}

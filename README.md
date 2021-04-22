## Description

## Paths table
Path | Method | Description | Body example
--- | --- | --- | ---
/courses | GET | Get all courses | [{"Code":207,"Title":"Mobile Application Development","DepartmentCode":5,"Description":"Mobile Application Development course description..."},{"Code":208,"Title":"Java Web Development","DepartmentCode":5,"Description":"Java Web Development course description..."},{"Code":209,"Title":"Architecture Operating Systems","DepartmentCode":5,"Description":"Architecture Operating Systems course description..."}]
/courses/{code} | GET | Get course by code | {"Code":207,"Title":"Mobile Application Development","DepartmentCode":5,"Description":"Mobile Application Development course description..."}
/courses | POST | Create new course |
/courses/{code} | PATCH | Update course description |
/courses/{code} | DELETE | Delete course by code |
## How to run  
1. Specify database and host data in the configuration file.
2. Run application with command: `go run main.go`
## Unit tests
```go test -race
```

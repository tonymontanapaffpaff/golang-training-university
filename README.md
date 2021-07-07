## Description
A small application representing CRUD operations on a university database.
## Paths table
<table>
<tr>
<td>Path</td>
<td>Method</td>
<td>Description</td>
<td>Body example</td>
</tr>
<tr>
<td>/courses</td>
<td>GET</td>
<td>Get all courses</td>
<td> 
```json
[
  {
    "Code":207,
    "Title":"Mobile Application Development",
    "DepartmentCode":5,
    "Description":"Mobile Application Development course description..."
  },
  {
    "Code":208,
    "Title":"Java Web Development",
    "DepartmentCode":5,
    "Description":"Java Web Development course description..."
  },
  {
    "Code":209,
    "Title":"Architecture Operating Systems",
    "DepartmentCode":5,
    "Description":"Architecture Operating Systems course description..."
  }
]
```
</td>
</tr>
<tr>
<td>/courses/{code}</td>
<td>GET</td>
<td>Get course by code</td>
<td>
  
```json
{
  "Code":207,
  "Title":"Mobile Application Development",
  "DepartmentCode":5,
  "Description":"Mobile Application Development course description..."
}
```
</td>
</tr>
<tr>
<td>/courses</td>
<td>POST</td>
<td>Create new course</td>
<td></td>
</tr>
<tr>
<td>/courses/{code}</td>
<td>PATCH</td>
<td>Update course description</td>
<td></td>
</tr>
<tr>
<td>/courses/{code}</td>
<td>DELETE</td>
<td>Delete course by code</td>
<td></td>
</tr>
</table>

## How to run  
1. Specify your database and host data in the configuration file: `/config/home_config.json`
2. Run application with command: `go run main.go`


## Run in a cluster

First you need to create a cluster:

```
k3d cluster create demo --servers 1 --agents 1 --port 8080:8080/TCP@loadbalancer
```

Then create all the necessary entities:
```
kubectl apply -f ./k3s/mongodb-config.yaml
```

The server will be aviable on port `8080`. 

## Unit tests
```
go test -race
```

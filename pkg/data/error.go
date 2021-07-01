package data

// HttpErr contains an error and status code of HTTP response
type HttpErr struct {
	Err        error
	StatusCode int
}

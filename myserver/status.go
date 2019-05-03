package main

type HttpStatus struct {
	Code    int
	Message string
}

func newHttpStatus(code int, message string) HttpStatus {
	res := HttpStatus{code, message}
	return res
}

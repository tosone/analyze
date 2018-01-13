package constants

type RequestType string

const (
	SummaryRequest RequestType = "summary"
)

type RequestMethod string

const (
	GET  RequestMethod = "GET"
	POST RequestMethod = "POST"
)

type RequestStatus string

const (
	Success RequestStatus = "SUCCESS"
	Error   RequestStatus = "ERROR"
)

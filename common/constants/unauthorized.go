package constants

type Unauthorized struct {
	Error string `json:"error"`
}

const UnauthorizedStr = "Unauthorized"

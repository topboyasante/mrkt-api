package types

type APIError struct {
	Message string `json:"error"`
}

type APISuccess struct {
	Message string `json:"success"`
	Data    any    `json:"data"`
}

package api

type Request struct {
	Operation string `json:"operation"`
	Number1   any    `json:"a"`
	Number2   any    `json:"b"`
}
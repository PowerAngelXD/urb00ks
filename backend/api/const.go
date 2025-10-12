package api

type Error struct {
	Error string `json:"error"`
}

var ErrorStruct = Error{Error: "A Error occurred"}

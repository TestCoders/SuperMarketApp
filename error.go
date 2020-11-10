package main

type Error struct {
	Type   string `json:"type"`
	Reason string `json:"reason"`
}

func NewError(err error) Error {
	return Error{
		Type:   "error",
		Reason: err.Error(),
	}
}

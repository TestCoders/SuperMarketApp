package main

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Type   string `json:"type"`
	Reason string `json:"reason"`
}

func newError(err error) Error {
	return Error{
		Type:   "error",
		Reason: err.Error(),
	}
}

func (a *application) serveError(w http.ResponseWriter, err error, statusCode int) {
	e := newError(err)
	w.WriteHeader(statusCode)

	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	if err = enc.Encode(e); err != nil {
		a.LogError(err)
	}
}

package views

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	view := json.NewEncoder(w)

	return view.Encode(data)
}

func JsonStatus(w http.ResponseWriter, data any, code int) {
	w.WriteHeader(code)
	Json(w, data)
}

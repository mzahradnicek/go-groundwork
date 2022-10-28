package views

import (
	"encoding/xml"
	"net/http"
)

func Xml(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-type", "text/xml; charset=utf-8")
	view := xml.NewEncoder(w)

	return view.Encode(data)
}

func XmlStatus(w http.ResponseWriter, data any, code int) {
	w.WriteHeader(code)
	Json(w, data)
}

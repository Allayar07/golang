package apiserver

import "net/http"

type REsponseWriter struct {
	http.ResponseWriter
	code int
}

func(w *REsponseWriter) WriteHEader(StatusCode int) {
	w.code = StatusCode
	w.ResponseWriter.WriteHeader(StatusCode)
}
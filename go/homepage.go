package main

import (
    "net/http"
    "net/http/cgi"

    "github.com/gorilla/mux"
)

// TODO: Move global constants into a singleton Config; read from JSON file
const (
    // Our subrouter URI on server
    basePath = "/~smn2028"
)

func main() {
    r := mux.NewRouter()
    // We are always under basePath
    s := r.PathPrefix(basePath).Subrouter()
    s.Handle("/testimonials", Testimonials{})
    s.Handle("/", Index{})
    http.Handle("/", r)
    cgi.Serve(nil)
}

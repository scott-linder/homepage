package main

import (
    "net/http/cgi"
    "log"

    "github.com/gorilla/mux"
)

// Cfg is the global config instance for the site.
var Cfg Config

func main() {
    if err := Cfg.Load("../config/global.json"); err != nil {
        log.Fatal(err)
    }
    r := mux.NewRouter()
    s := r.PathPrefix(Cfg.BasePath).Subrouter()
    s.Handle("/testimonials", Testimonials{})
    s.Handle("/", Index{})
    cgi.Serve(r)
}


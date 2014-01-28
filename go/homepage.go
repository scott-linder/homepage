package main

import (
    "net/http/cgi"
    "log"

    "blog"

    "github.com/gorilla/mux"
)

// Cfg is the global config instance for the site.
var Cfg Config

func main() {
    if err := Cfg.Load("../config/global.json"); err != nil {
        log.Fatal(err)
    }
    r := mux.NewRouter().StrictSlash(false)
    s := r.PathPrefix(Cfg.BasePath).Subrouter()
    st := s.PathPrefix("/testimonials").Subrouter()
    blog.NewBlog(st, "../template/testimonials.tpl", "../testimonials/", Cfg.PageSize)
    s.Handle("/", Index{})
    cgi.Serve(r)
}


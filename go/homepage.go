package main

import (
    "net/http/cgi"
    "log"

    "github.com/scott-linder/blog"
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
    blog.NewBlogSimple("testimonials", st)
    sb := s.PathPrefix("/blag").Subrouter()
    blog.NewBlogSimple("blag", sb)
    s.Handle("/", Index{})
    cgi.Serve(r)
}



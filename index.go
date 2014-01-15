package main

import (
    "log"
    "net/http"
    "html/template"
)

type Index struct {}

func (self Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    tpl, err := template.ParseFiles("../template/index.tpl")
    if err != nil {
        log.Fatal(err.Error())
    }

    data := struct {
    } {}

    tpl.Execute(w, data)
}

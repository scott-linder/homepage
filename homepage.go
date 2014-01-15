package main

import (
    "fmt"
    "net/http"
    "net/http/cgi"

    "github.com/gorilla/mux"
)

// TODO: Move global constants into a singleton Config; read from JSON file
const (
    // Our subrouter URI on server
    basePath = "/~smn2028"
    // Temporary page
    html = `
<!DOCTYPE html>
<html>
<head>
    <title>It's Nothing</title>
    <meta charset='utf-8'>
    <style type='text/css'>
        body {
            background: #2c2c2c;
            color: #dcdccc;
        }
        section {
            width: 500px;
            margin: auto;
        }
    </style>
</head>
<body>
    <section>
        <h1>Hompage of Scott Linder</h1>
        <h2>Testimonials</h2>
        <blockquote>
        <p>Scott linder, world leader, capture the flagger, and an overall
        pretty cool guy.</p>
        <cite>Mrlrl</cite>
        </blockquote>
        <h2>Dew</h2>
        <img src="http://www.mountaindew.com/assets/50c49j.png" alt="Dew"
            height="590" width="390">
    </section>
</body>
</html>
`
)

func main() {
    r := mux.NewRouter()
    // We are always under basePath
    s := r.PathPrefix(basePath).Subrouter()
    s.Handle("/testimonials", Testimonials{})
    s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, html)
    })
    http.Handle("/", r)
    cgi.Serve(nil)
}

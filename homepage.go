package main

import (
    "fmt"
    "net/http"
    "net/http/cgi"
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
        <p>Or something.</p>
    </section>
</body>
</html>
`
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, html)
    })
    cgi.Serve(nil)
}

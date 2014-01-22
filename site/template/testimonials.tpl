<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
    <meta charset='utf-8'>
    <link rel="stylesheet" type="text/css" href="/~smn2028/style/main.css">
</head>
<body>
    <section>
        <h1>Testimonials</h1>
        <h2>Where others blog vicariously through quotes.</h2>
    </section>
    <section>
        {{range $testimonial := .Testimonials}}
        <article>
        {{$testimonial.Body}}
        </article>
        {{else}}
        <p>There ain't no testimonials.</p>
        {{end}}
    </section>
    <section>
        <footer>
            {{.Footer}}
        </footer>
    </section>
</body>
</html>

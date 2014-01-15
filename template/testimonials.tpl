<!DOCTYPE html>
<html>
<head>
    <title>Testimonials</title>
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
        <dl>
        <dt>Last Modified</dt>
        <dd>{{$testimonial.ModTime}}</dd>
        </dl>
        </article>
        {{else}}
        <p>There ain't no testimonials.</p>
        {{end}}
    </section>
    <section>
        <footer>
        © Scott Linder (2014)
        </footer>
    </section>
</body>
</html>

<!DOCTYPE html>
<html>
<head>
    <title>Scott Linder − Testimonials</title>
    <meta charset='utf-8'>
    <link rel="stylesheet" type="text/css" href="/~smn2028/style/main.css">
</head>
<body>
    <section>
        <h1>Testimonials</h1>
        <h2>Where others blog vicariously through quotes.</h2>
    </section>
    <section>
        {{range $post := .Posts}}
        <article>
        <time datetime="{{$post.Date.Format "2006-01-02"}}">
            {{$post.Date.Format "Jan 2, 2006"}}</time>
        <h1>{{$post.Name}}</h1>
        {{$post.Body}}
        <a href="{{$post.Permalink}}" class="permalink">Permalink</a>
        </article>
        {{else}}
        <p>There ain't no testimonials.</p>
        {{end}}
    </section>
</body>
</html>

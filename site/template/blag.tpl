<!DOCTYPE html>
<html>
<head>
    <title>Scott Linder − Blag</title>
    <meta charset='utf-8'>
    <link rel="stylesheet" type="text/css" href="/~smn2028/style/main.css">
    <script src="https://google-code-prettify.googlecode.com/svn/loader/run_prettify.js?lang=go&amp;skin=sunburst"></script>
</head>
<body>
    <section>
        <h1>Blag</h1>
        <h2>n. A weblag.</h2>
    </section>
    <section>
        {{range $post := .Posts}}
        <article>
        <time datetime="{{$post.Date.Format "2006-01-02"}}">
            {{$post.Date.Format "Jan 2, 2006"}}</time>
        <h1>{{$post.Name}}</h1>
        {{$post.Body}}
        <a href="{{$post.Permalink}}" class="permalink">Immutalink</a>
        </article>
        {{else}}
        <p>There ain't no blags.</p>
        {{end}}
    </section>
</body>
</html>

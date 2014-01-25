<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
    <meta charset='utf-8'>
    <link rel="stylesheet" type="text/css" href="/~smn2028/style/main.css">
</head>
<body>
    <section>
        <h1>Homepage of Scott Linder</h1>
        <p>Feel free to look around and not touch anything.</p>
    </section>
    <section>
        <h2>About</h2>
        <h3>Me</h3>
        <p>I'm currently attending Western Michigan University, pursuing a degree
        in Computer Science with a possible minor in Chemistry. I like programming,
        operating systems, computers, and Oxford commas.</p>
        <h3>The Site</h3>
        <p>This is my homepage on the WMU CS webservers (shout out to the
        department for free hosting). The site is largely written in Go, and
        all of the sources are avaialable on my GitHub at
        <a href="https://github.com/scott-linder/homepage">scott-linder/homepage</a>.
        </p>
        <h2>Navigation</h2>
        <p>Get to the myriad (read: <em>one</em>) parts of the site.</p>
        <ul>
            <li><a href="testimonials">Testimonials</a></li>
        </ul>
        <h2>Related</h2>
        <p>Stuff related to my existence.</p>
        <ul>
            <li><a href="https://github.com/scott-linder">My Github</a></li>
            <li><a href="http://ccowmu.org/">WMU Computer Club</a></li>
        </ul>
    </section>
    <section>
        <footer>
            {{.Footer}}
        </footer>
    </section>
</body>
</html>

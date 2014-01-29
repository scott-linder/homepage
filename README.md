# Homepage of Scott Linder #

My [homepage](http://cs.wmich.edu/~smn2028) at the WMU CS department.

## Rationale ##

The site is written in [Go](http://golang.org/) as an exercise (the site is
could easily be completely static). Since it is an academic site which is in
git anyway I'm just putting the sources up here in case anyone is
interested.

## Structure ##
Things are divided up into a few directories and files to keep things sane and
make it easier to automatically build/deploy.

### / ###
Site meta-data, build scripts, deploy scripts, etc.

### go/ ###
Source for the dynamic portion (CGI binary) of the website.

### site/ ###
Anything which the website runs from but which should not be served publicly
on the server; examples include raw blog entries.

#### template/ ###
Most Go source files have an associated HTML template file.

#### config/ ###
Mostly json config file(s).

### site/public/ ###
Anything which the web server can serve publicly.


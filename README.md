
# Homepage of Scott Linder

## Structure
Things are divided up into a few directories and files to keep things sane and
make it easier to automatically build/deploy.

### /
Site meta-data, build scripts, deploy scripts, etc.

#### makefile
The site is built and "packaged" by GNU make.

#### deploy.sh
Small shell script to deploy built site.

### go/
Source for the dynamic portion (CGI binary) of the website.

#### homepage.go
The entry point for the CGI website. Try to keep it clean and simple.

#### *.go
Most Go source files encapsulate a complete page or group of pages on the site.

### private/
Anything which the website runs from but which should not be served publicly
on the server; examples include raw blog entries.

#### template/
Most Go source files have an associated HTML template file.

### public/
Anything which the web server can serve publicly.

### homepage/
This directory is where all `make` output is housed, in the correct format
to be sync'd directly to the web server by deploy.sh


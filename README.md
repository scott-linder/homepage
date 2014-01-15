
# Homepage of Scott Linder

## Structure
Things are divided up into a few directories and files to keep things sane and
make it easier to automatically build/deploy.

### makefile
The site is built and "packaged" by GNU make.

### deploy.sh
Small shell script to deploy built site.

### homepage
This directory is where all `make` output is housed, in the correct format
to be sync'd directly to the web server by deploy.sh

### public
Anything in this directory represents content to be served as-is by the web
server. It will also contain the final cgi binary.

### homepage.go
The entry point for the cgi website. Try to keep it clean and simple.

### *.go
Most Go source files encapsulate a complete page or group of pages on the site.

### *.tpl
Most GO source files have an associated HTML template file.

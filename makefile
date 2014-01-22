#
# Homepage makefile.
# Authors:
# 	Scott Linder
#

# Private data output directory.
OUT=site/
# Public data output directory.
PUB:=$(OUT)public/

# Go source.
GO_SRC:=$(shell find go/ -type f -name '*.go')
# Go CGI.
CGI=$(PUB)router.cgi


# Build site.
.PHONY: all
all: $(CGI)

# Build site CGI.
$(CGI): $(GO_SRC)
	cd go; go build -o ../$(CGI)

.PHONY: clean
clean:
	-rm $(CGI)


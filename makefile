#
# Homepage makefile.
# Authors:
# 	Scott Linder
#

# Site output directory.
OUT=homepage/

# Site input directories.
PRI=private/
PUB=public/

# Source files for public and private parts of site.
PRI_SRC:=$(shell find $(PRI) -type f)
PUB_SRC:=$(shell find $(PUB) -type f)
# Destinations of source files.
PRI_DEST:=$(patsubst private/%, $(OUT)%, $(PRI_SRC))
PUB_DEST:=$(addprefix $(OUT), $(PUB_SRC))
# Combine all parts of site so we can have one copy rule.
FILES_SRC:=$(PRI_SRC) $(PUB_SRC)
FILES_DEST:=$(PRI_DEST) $(PUB_DEST)

# Go source.
GO_SRC:=$(shell find go/ -type f -name '*.go')
# Go CGI.
CGI=$(OUT)$(PUB)router.cgi


# Build site.
all: $(FILES_DEST) $(CGI)

$(OUT)%: $(PRI)%
	@mkdir -p $(@D)
	cp $< $@

$(OUT)$(PUB)%: $(PUB)%
	@mkdir -p $(@D)
	cp $< $@

# Build site CGI.
$(CGI): $(GO_SRC)
	cd go; go build -o ../$(CGI)

clean:
	-rm -rf $(OUT)


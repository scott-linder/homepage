#
# Homepage makefile
# Authors:
# 	Scott Linder
#

IN=./
OUT=./homepage/

PUB=$(OUT)public/
CGI=$(PUB)router.cgi

all: $(CGI)

$(CGI): $(PUB)
	go build -o $(CGI)

$(PUB): | $(OUT)
	cp -r $(IN)public/ $(PUB)

$(OUT):
	[ -d $(OUT) ] || mkdir $(OUT)

clean:
	-rm -rf $(OUT)


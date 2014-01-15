#
# Homepage makefile
# Authors:
# 	Scott Linder
#

IN=./
OUT=./homepage/

PUB=$(OUT)public/
TPL=$(OUT)template/
TST=$(OUT)testimonials/
CGI=$(PUB)router.cgi

all: $(CGI) $(TPL) $(TST)

$(TST):
	cp -r $(IN)testimonials/ $(TST)

$(TPL):
	cp -r $(IN)template/ $(TPL)

$(CGI): $(PUB)
	go build -o $(CGI)

$(PUB): | $(OUT)
	cp -r $(IN)public/ $(PUB)

$(OUT):
	[ -d $(OUT) ] || mkdir $(OUT)

clean:
	-rm -rf $(OUT)


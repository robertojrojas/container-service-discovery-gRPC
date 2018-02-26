SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

PRJ_NAME=greeter-go
BINARY=${PRJ_NAME}.bin
VERSION=1.0.2
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-a -tags netgo -ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -w" 

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	CGO_ENABLED=0 GOOS=linux go build ${LDFLAGS} -o ${BINARY} *.go

docker: clean ${BINARY}
	docker build -t robertojrojas/${PRJ_NAME}:${VERSION} .

docker_clean: clean
	docker rmi robertojrojas/${PRJ_NAME}:${VERSION}

.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

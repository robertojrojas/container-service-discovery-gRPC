SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

BINARY=greeter-go
VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-a -tags netgo -ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -w" 

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	CGO_ENABLED=0 GOOS=linux go build ${LDFLAGS} -o ${BINARY} *.go

docker: clean ${BINARY}
	docker build -t robertojrojas/${BINARY}:${VERSION} .

docker_clean: clean
	docker rmi robertojrojas/${BINARY}:${VERSION}

.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
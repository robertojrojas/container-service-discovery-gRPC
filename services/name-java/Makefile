SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.java')

BINARY=name-java
VERSION=1.0.0

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	mvn clean package

docker: clean ${BINARY}
	docker build -t robertojrojas/${BINARY}:${VERSION} .

docker_clean: clean
	docker rmi robertojrojas/${BINARY}:${VERSION} 


.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

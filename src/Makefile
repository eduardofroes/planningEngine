GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=${GOCLEAN} clean
GOGET=${GOCMD} get
GOMOD=${GOCMD} mod
GOMODWHY=${GOMOD} why
GOMODINIT=${GOMOD} init
BINARY_NAME=../bin/planningEngine

build:
	${GOBUILD} -o $(BINARY_NAME) -v 

clean:
	${GOCLEAN}

why:
	${GOMODWHY}

run: build
	${BINARY_NAME}
	
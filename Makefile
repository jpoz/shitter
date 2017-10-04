BINARY=shitter

VERSION=1.0.0

REPO=github.com/jpoz/shitter

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '**/*.go')

.DEFAULT_GOAL: $(BINARY)

$(BINARY): $(SOURCES)
	go build -o ${BINARY} cmd/shitter/main.go

.PHONY: install
install:
	go install ${LDFLAGS} ./...

.PHONY: clean
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: assets
assets:
	go-bindata -pkg shitter shit.png

.PHONY: run
run: clean $(BINARY)
	./$(BINARY) test.png out.png

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=magic_arena_game

build:
	$(GOBUILD) -o out/$(BINARY_NAME) -v
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run: build
	./out/$(BINARY_NAME)
deps:
	$(GOGET) github.com/stretchr/testify/assert
test :
	go test ./... -v | grep RUN -v
test.cover:
	go test ./... -cover
GO=go
BIN=bin
SRC=main.go processors.go
EXEC=customwc

all: build

build: $(SRC)
	$(GO) build -o $(BIN)/$(EXEC) $(SRC)

test:
	$(GO) test

clean:
	rm -f bin/*
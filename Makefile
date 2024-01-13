GO=go
BIN=bin
EXEC=customwc

build:
	$(GO) build -o $(BIN)/$(EXEC) main.go

clean:
	rm -f bin/*
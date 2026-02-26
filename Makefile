.PHONY: run build clean help

help:
    @echo "rocketGo - Makefile Commands"
    @echo "=============================="
    @echo "make run      - Run the game"
    @echo "make build    - Build executable"
    @echo "make clean    - Remove build artifacts"

run:
    go run main.go

build:
    go build -o rocketGo.exe main.go

clean:
    rm -f rocketGo.exe
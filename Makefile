# Makefile
.PHONY: all

all: build run

build:
    go build main.go

run:
    ./main.exe todo.go
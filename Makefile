ifeq ($(OS),Windows_NT)
	PROC_NAME=urlshortener.exe
	BINARY=.\bin\urlshortener.exe
	KILL = taskkill /F /IM $(PROC_NAME)
else
	PROC_NAME=urlshortener
	BINARY=./bin/urlshortener
	KILL = pkill -f $(BINARY)
endif

build:
	go build -o $(BINARY) ./cmd/app

run:
	start /B $(BINARY) &

stop:
	$(KILL)

start: build run

restart: stop build run

ifeq ($(OS),Windows_NT)
	BINARY=./bin/urlshortener.exe
    OSFLAG=WINDOWS
else
	BINARY=./bin/urlshortener
    OSFLAG=UNIX
endif

ifeq ($(OSFLAG),WINDOWS)
	KILL = taskkill /F /IM $(BINARY)
else
	KILL = pkill -f $(BINARY)
endif

build:
	go build -o $(BINARY) ./cmd/app

run:
	$(BINARY) &

stop:
	$(KILL)


start: build run

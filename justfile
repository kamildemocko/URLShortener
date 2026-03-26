[linux]
build:
    go build -o ./bin/urlshortener ./cmd/urlshortener

[linux]
run: build
    ./bin/urlshortener &

[windows]
run: build
    start /B ./bin/urlshortener.exe &

[linux]
stop:
    -pkill -f ./bin/urlshortener

[windows]
stop:
    -taskkill /F /IM urlbinshortener

[linux]
open:
    xdg-open http://localhost:8080 &

[windows]
open:
    open http://localhost:8080

generate:
    @templ generate

start: stop generate build run open

restart: stop generate build run
# URL Shortener

This is a simple URL shortener application built with Go and Docker. App can be run locally or in Docker.

## Prerequisites

- Go
- Docker
- Postgres
    - used schema is `urlshortener`
    - table `keys` is created automatically
- [opt] Make (for running Makefile)
- [opt] Bruno (for endpoint tests)

## Installation

### .env File Preparation

1. Create a new file `.env` in the root directory of the project
2. Add the following environment variables to the file:

```
DSN=host=192.168.88.221 port=5432 user=postgres password=SECRET sslmode=disable timezone=UTC connect_timeout=5 search_path=uslshortener
PORT=80
DOMAIN=localhost
PROTOCOL=http
```

### Docker

1. Clone the repository

```
git clone https://github.com/kamildemocko/URLShortener.git
```

2. Create network and make sure DB server has access to the same network

```
docker network url-shortener-net
```

3. Build the Docker image

```
docker-compose build
```

4. Run the Docker container

```
docker-compose up -d
```

### Local

1. Clone the repository

```
git clone https://github.com/kamildemocko/URLShortener.git
```

2. Download dependencies

```
go mod download
```

3. Run 

```
make start
```

## Usage

1. Open your web browser and navigate to `http://localhost:PORT/short` (where PORT is port that is set up in .env file or docker-compose.yaml)
2. Enter a custom short link key and a long URL
3. Click on "Get short link"
4. Your short link will be displayed
5. Click on "Copy to Clipboard" to copy the short link

## Example

[https://www.hrasok.xyz/short](https://www.hrasok.xyz/short)


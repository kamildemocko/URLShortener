# URL Shortener

This is a simple URL shortener application built with Go and Docker.

## Prerequisites

- Docker
- Go

## Installation

1. Clone the repository

```
git clone https://github.com/kamildemocko/URLShortener.git
```

2. Build the Docker image

```
docker-compose build
```

3. Run the Docker container

```
docker-compose up -d
```

## Usage

1. Open your web browser and navigate to `http://localhost:80`
2. Enter a custom short link key and a long URL
3. Click on "Get short link"
4. Your short link will be displayed
5. Click on "Copy to Clipboard" to copy the short link

# Notification Listener

## Setup

1. Install [Go](https://go.dev/dl/).
2. Copy `.env.example` to `.env` and set the required values.
3. Ensure you have the public Apple certificate set in `.env`.

## Running the Server

```sh
make dev
```

The server will be available at http://localhost:8080.

## Running Tests

```sh
make test
```

## Code Generation

```sh
make generate
```

## API

**POST** `/v1/apple-notifications`

## Simple Envoy Setup

### Steps

1. Run the Docker Compose setup

```sh
docker compose up -d
docker compose logs -f
```

2. Invoke the service

```sh
curl -v http://localhost:8000/foo/baz
```

```sh
curl -v http://localhost:8000/baz
```
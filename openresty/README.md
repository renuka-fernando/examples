# OpenResty Example

## Run Docker Compose Setup

```sh
docker compose down; docker compose up -d; docker compose logs -f
```

## Sample cURLs

```sh
curl 'http://localhost:8080?pretty=true' -H 'Authorization: Bearer fooo' -i
curl 'http://localhost:8080/echo?delayMs=5000' -i
curl 'http://localhost:8080/nginx-health' -i
curl 'http://localhost:8080/foo' -i
```

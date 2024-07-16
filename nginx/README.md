## Nginx

### Run the Setup

```sh
docker compose down; docker compose up -d; docker compose logs -f
```

### Test
```sh
curl 'http://localhost:8080/foo?foo=bar&api_key=1234&abc=baz' -i \
    -H 'API-Key: abcd-bar-baz'
```

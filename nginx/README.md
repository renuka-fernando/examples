## Nginx

### Run the Setup

```sh
docker compose down; docker compose up -d; docker compose logs -f
```

### Test - Nginx
```sh
curl 'http://localhost:8080/foo?foo=bar&api_key=1234&abc=baz' -i \
    -H 'API-Key: abcd-bar-baz'
```

### Test - OpenResty
```sh
curl 'http://localhost:18080/foo?foo=bar&api_key=1234&abc=baz' -i \
    -H 'API-Key: abcd-bar-baz'
```

## Nginx

### Run the Setup

```sh
docker compose down; docker compose up -d; docker compose logs -f
```

### Test - Envoy
```sh
curl 'http://localhost:8000/foo?foo=bar&api_key=1234&abc=baz' -i \
    -H 'API-Key: abcd-bar-baz'
```

### Test - OpenResty

#### HTTP

```sh
curl 'http://localhost:18080/foo?foo=bar&api_key=1234&abc=baz' -i \
    -H 'API-Key: abcd-bar-baz'
```

#### mTLS

```sh
curl 'https://localhost:18443/foo?foo=bar&api_key=1234&abc=baz' -i \
    -H 'API-Key: abcd-bar-baz' --cacert certs/listener.pem --cert certs/client.pem --key certs/client.key
```


### Test - Nginx
```sh
curl 'http://localhost:28080/foo?foo=bar&api_key=1234&abc=baz' -i \
    -H 'API-Key: abcd-bar-baz'
```

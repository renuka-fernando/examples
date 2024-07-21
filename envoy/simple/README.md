## Simple Envoy Setup

### Steps

1. Run the Docker Compose setup

```sh
docker compose down; docker compose up -d; docker compose logs -f -t
```

2. Invoke the service

```sh
curl -v 'http://localhost:8000/foo/baz'
```

```sh
curl -v 'http://localhost:8000/baz'
```

## Scenarios

Access log format:
[%START_TIME%] "%REQ(:METHOD)% %REQ(X-ENVOY-ORIGINAL-PATH?:PATH)% %PROTOCOL%" %RESPONSE_CODE% %RESPONSE_FLAGS% %RESPONSE_CODE_DETAILS% %BYTES_RECEIVED% %BYTES_SENT% %DURATION% %RESP(X-ENVOY-UPSTREAM-SERVICE-TIME)% "%REQ(X-FORWARDED-FOR)%" "%REQ(USER-AGENT)%" "%REQ(X-REQUEST-ID)%" "%REQ(:AUTHORITY)%" "%UPSTREAM_HOST%"

### UT: HTTP 502: Gateway Timeout

```sh
curl -v 'http://localhost:8000/gateway-timeout?delayMs=5000'
```

```log
request_info-1  | 2024-07-20T07:05:54.173886083Z 2024/07/20 07:05:54 [INFO] "GET" "/gateway-timeout?delayMs=5000", Host: "localhost:8000", Content Length: 0, "curl/8.4.0", "172.27.0.3:52020"
envoy-1         | 2024-07-20T07:05:58.670998585Z [2024-07-20T07:05:54.173Z] "GET /gateway-timeout?delayMs=5000 HTTP/1.1" 504 UT response_timeout 0 24 4497 - "-" "curl/8.4.0" "6d67d674-2122-4a2d-b172-0be70c674371" "localhost:8000" "172.27.0.2:8080"
```

### UH: HTTP 503 Service Unavailable: No Healthy Upstream

```sh
curl -v 'http://localhost:8000/no-healthy-upstream'
```

```log
envoy-1         | 2024-07-20T07:06:24.756083556Z [2024-07-20T07:06:24.729Z] "GET /no-healthy-upstream HTTP/1.1" 503 UH no_healthy_upstream 0 19 0 - "-" "curl/8.4.0" "2f22e5c6-2387-49b8-96c4-7b3a1678dbb0" "localhost:8000" "-"
```


### UF: HTTP 503 Service Unavailable: Upstream Connection Failure

```sh
curl -v 'http://localhost:8000/upstream-connection-failure'
```

```log
envoy-1         | 2024-07-20T07:09:48.471209278Z [2024-07-20T07:09:48.470Z] "GET /upstream-connection-failure HTTP/1.1" 503 UF upstream_reset_before_response_started{connection_failure,delayed_connect_error:_111} 0 145 0 - "-" "curl/8.4.0" "44898dba-8d8e-48f4-93f7-23e4b4c592c0" "localhost:8000" "172.28.0.2:801"
```

### UPE: HTTP 200: Upstream Protocol Error

```sh
python3 unstable_backend_server.py
```

```sh
curl -v 'http://localhost:8000/upstream-connection-failure/after-response'
```

```log
envoy-1         | 2024-07-20T07:23:59.593171059Z [2024-07-20T07:23:58.586Z] "GET /upstream-connection-failure/after-response HTTP/1.1" 200 UPE upstream_reset_after_response_started{protocol_error} 0 19 1006 0 "-" "curl/8.4.0" "84ef2354-e85c-4e90-bc6e-bb81d5dc70e8" "localhost:8000" "192.168.5.2:8080"
```


### UO: HTTP 503 Service Unavailable: Upstream Overflow

```sh
curl -v 'http://localhost:8000/upstream-service-overflow?delayMs=5000' & \
    curl -v 'http://localhost:8000/upstream-service-overflow?delayMs=5000' & \
    curl -v 'http://localhost:8000/upstream-service-overflow?delayMs=5000'
```

```sh
jmeter -t test.jmx \
-Jusers=500 \
-JrampUpPeriod=10 \
-Jduration=1020 \
-Jproto=http \
-Jport=8000 \
-Jip='localhost' \
-Jpath='/upstream-service-overflow?delayMs=5000' \
-Jpayload="/Users/renuka/Documents/payloads/50B.json" \
-n -l results.jtl
```

```log
request_info-1  | 2024-07-20T07:32:21.516940173Z 2024/07/20 07:32:21 [INFO] "GET" "/upstream-service-overflow?delayMs=5000", Host: "localhost:8000", Content Length: 0, "curl/8.4.0", "172.31.0.2:44442"
envoy-1         | 2024-07-20T07:32:21.520419507Z [2024-07-20T07:32:21.519Z] "GET /upstream-service-overflow?delayMs=5000 HTTP/1.1" 503 UO upstream_reset_before_response_started{overflow} 0 81 0 - "-" "curl/8.4.0" "c414b70a-7c85-4f78-9c35-2ce2b9131aa3" "localhost:8000" "172.31.0.3:8080"
envoy-1         | 2024-07-20T07:32:26.551823467Z [2024-07-20T07:32:21.516Z] "GET /upstream-service-overflow?delayMs=5000 HTTP/1.1" 200 - via_upstream 0 682 5002 5001 "-" "curl/8.4.0" "dcf9d063-610a-4f11-96db-db890542d64a" "localhost:8000" "172.31.0.3:8080"
```

### NC: HTTP 503 Service Unavailable: No Cluster Found

```sh
curl -v 'http://localhost:8000/no-upstream-cluster' -H 'x-envoy-upstream-cluster: no-upstream-cluster'
```

```log
envoy-1         | 2024-07-20T07:38:15.900226551Z [2024-07-20T07:38:15.899Z] "GET /no-upstream-cluster HTTP/1.1" 503 NC cluster_not_found 0 0 0 - "-" "curl/8.4.0" "8a23b89d-1dcd-4f5c-ae55-58e79d74d8ca" "localhost:8000" "-"
```

### URX: HTTP 5xx: Upstream Retry Limit Exceeded

```sh
curl -v 'http://localhost:8000/upstream-retry-exceeds?statusCode=500'
```

```log
request_info-1  | 2024-07-20T07:45:34.829817885Z 2024/07/20 07:45:34 [INFO] "GET" "/upstream-retry-exceeds?statusCode=500", Host: "localhost:8000", Content Length: 0, "curl/8.4.0", "192.168.48.3:57404"
request_info-1  | 2024-07-20T07:45:34.837428760Z 2024/07/20 07:45:34 [INFO] "GET" "/upstream-retry-exceeds?statusCode=500", Host: "localhost:8000", Content Length: 0, "curl/8.4.0", "192.168.48.3:50228"
request_info-1  | 2024-07-20T07:45:34.879198385Z 2024/07/20 07:45:34 [INFO] "GET" "/upstream-retry-exceeds?statusCode=500", Host: "localhost:8000", Content Length: 0, "curl/8.4.0", "192.168.48.3:50244"
request_info-1  | 2024-07-20T07:45:34.887870218Z 2024/07/20 07:45:34 [INFO] "GET" "/upstream-retry-exceeds?statusCode=500", Host: "localhost:8000", Content Length: 0, "curl/8.4.0", "192.168.48.3:50260"
envoy-1         | 2024-07-20T07:45:34.917997010Z [2024-07-20T07:45:34.828Z] "GET /upstream-retry-exceeds?statusCode=500 HTTP/1.1" 500 URX via_upstream 0 680 59 59 "-" "curl/8.4.0" "0df38f52-08eb-4b67-a097-ece7427c7238" "localhost:8000" "192.168.48.2:8080"
```

## WebSocket

### Start server

```sh
node websocket-server-nodejs/server.js
```

### Tryout

```sh
websocat ws://localhost:8000/ws
```

### Create 100 connections

```sh
for i in {1..100}; do; 
    tail -f foo.txt | websocat ws://localhost:8000/ws & \
        sleep 0.1
done
```

```sh
for i in {1..100}; do; 
    echo 12345678 >> foo.txt
    sleep 1
done
```




```log
envoy-1           | 2024-07-21T08:33:12.906672966Z [2024-07-21T08:32:59.547Z] "GET /ws HTTP/1.1" 101 DC downstream_remote_disconnect 61 105 13342 - "-" "-" "e5248829-b318-443a-a763-e2460530f070" "localhost:8000" "192.168.5.2:8080"

envoy-1           | 2024-07-21T09:18:26.815973385Z [2024-07-21T09:11:33.481Z] "GET /ws HTTP/1.1" 101 SI stream_idle_timeout 100 154 413322 - "-" "-" "8a458ce2-bb8d-4e80-a305-f323bbab338f" "localhost:8000" "192.168.5.2:8080"
```

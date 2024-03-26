## Rate Limit

### Steps

1.  Start Docker Compose
    ```bash
    docker compose up -d; docker compose logs -f
    ```
2.  Invoke the following curl (3 times) to test the rate limit **2 Per Min**.
    ```bash
    curl http://localhost:8888/header -H "foo: foo" -i
    echo ""
    curl http://localhost:8888/header -H "foo: foo" -i
    echo ""
    curl http://localhost:8888/header -H "foo: foo" -i
    ```

    The output of the final command should be:
    ```bash
    HTTP/1.1 429 Too Many Requests
    ```

3.  Stop the Docker Compose
    ```bash
    docker compose down
    ```

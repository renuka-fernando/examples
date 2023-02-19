## Lua Filter

### Steps

Execute following commands from this directory.

1.  Run docker compose setup.
    ```
    docker compose up -d; docker compose logs -f
    ```
2.  Make requests.
    ```sh
    curl localhost:9090/multiple/lua/scripts -i
    ```

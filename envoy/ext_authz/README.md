## Ext Authz

### Steps

Execute following commands from this directory.

1.  Do changes and run ext-authz service in host machine.
    ```sh
    cd auth/grpc-service; go run main.go; cd -
    ```
2.  Generate certs.
    ```sh
    ./gen-certs.sh
    ```
3.  Run docker compose setup.
    ```
    docker compose up -d; docker compose logs -f
    ```
4.  Make requests.
    ```
    ./curl.sh
    ```

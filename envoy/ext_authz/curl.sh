#!/bin/bash

curl https://localhost:8000 -i -H 'Authorization: Bearer token1' \
  --cert certs/client.pem --key certs/client.key --cacert certs/listener.pem

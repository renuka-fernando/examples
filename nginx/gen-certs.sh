#!/bin/bash

rm -rf certs
mkdir -p certs
cd certs

CN="localhost"      
SAN="DNS:localhost, DNS:host.docker.internal"

generate_certs() {
    FILE_NAME="$1"
    echo $FILE_NAME
    openssl req -x509 -nodes -newkey rsa:4096 -keyout "${FILE_NAME}.key" -out "${FILE_NAME}.pem" -sha256 -subj "/CN=${CN}" -reqexts SAN -extensions SAN \
    -days 3650 \
    -subj "/CN=${CN}/C=LK/ST=CO/L=Colombo/O=Example/OU=Example" \
    -config <(cat /etc/ssl/openssl.cnf <(printf "[SAN]\nsubjectAltName=${SAN}"))
}

generate_certs client
generate_certs listener

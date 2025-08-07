#!/bin/bash

# Generate self-signed cert with SAN for my-foo.com

DOMAIN="my-foo.com"
OUTPUT_DIR="./certs"
DAYS=365

# Create output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Generate private key
openssl genrsa -out "$OUTPUT_DIR/$DOMAIN.key" 2048

# Create CSR config file with SAN
cat > "$OUTPUT_DIR/csr.conf" <<EOF
[req]
default_bits = 2048
prompt = no
default_md = sha256
distinguished_name = dn
req_extensions = req_ext

[dn]
C = US
ST = California
L = San Francisco
O = MyOrg
OU = Dev
CN = $DOMAIN

[req_ext]
subjectAltName = @alt_names

[alt_names]
DNS.1 = $DOMAIN
DNS.2 = *.$DOMAIN
EOF

# Generate CSR
openssl req -new -key "$OUTPUT_DIR/$DOMAIN.key" -out "$OUTPUT_DIR/$DOMAIN.csr" -config "$OUTPUT_DIR/csr.conf"

# Create cert config file for SAN
cat > "$OUTPUT_DIR/cert.conf" <<EOF
authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
subjectAltName = @alt_names

[alt_names]
DNS.1 = $DOMAIN
DNS.2 = *.$DOMAIN
EOF

# Generate self-signed certificate
openssl x509 -req \
  -in "$OUTPUT_DIR/$DOMAIN.csr" \
  -signkey "$OUTPUT_DIR/$DOMAIN.key" \
  -out "$OUTPUT_DIR/$DOMAIN.crt" \
  -days "$DAYS" \
  -sha256 \
  -extfile "$OUTPUT_DIR/cert.conf"

# Clean up temporary files
rm "$OUTPUT_DIR/csr.conf" "$OUTPUT_DIR/cert.conf" "$OUTPUT_DIR/$DOMAIN.csr"

echo "Certificate generated successfully:"
echo " - Private key: $OUTPUT_DIR/$DOMAIN.key"
echo " - Certificate: $OUTPUT_DIR/$DOMAIN.crt"
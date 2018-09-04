#!/usr/bin/env bash

# Create CSR
openssl req \
-config intermediate/openssl.cnf \
-new -sha256 \
-key intermediate/private/intermediate.key.pem \
-out intermediate/csr/intermediate.csr.pem

# Sign and create certificate
rm -f intermediate/certs/intermediate.cert.pem
openssl ca -config openssl.cnf -extensions v3_intermediate_ca \
-days 3650 -notext -md sha256 \
-in intermediate/csr/intermediate.csr.pem \
-out intermediate/certs/intermediate.cert.pem
chmod 444 intermediate/certs/intermediate.cert.pem
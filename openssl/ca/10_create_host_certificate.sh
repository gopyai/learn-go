#!/usr/bin/env bash

HOST=api.stefarf.com
EXT=server_cert
#EXT=usr_cert

# Create CSR
openssl req -config intermediate/openssl.cnf \
-key intermediate/private/$HOST.key.pem \
-new -sha256 -out intermediate/csr/$HOST.csr.pem

# Sign and create certificate
rm -f intermediate/certs/$HOST.cert.pem
openssl ca -config intermediate/openssl.cnf \
-extensions $EXT -days 375 -notext -md sha256 \
-in intermediate/csr/$HOST.csr.pem \
-out intermediate/certs/$HOST.cert.pem
chmod 444 intermediate/certs/$HOST.cert.pem
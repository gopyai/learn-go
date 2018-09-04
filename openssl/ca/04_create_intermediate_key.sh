#!/usr/bin/env bash
rm -f intermediate/private/intermediate.key.pem
openssl genrsa -aes256 \
-out intermediate/private/intermediate.key.pem 4096
chmod 400 intermediate/private/intermediate.key.pem
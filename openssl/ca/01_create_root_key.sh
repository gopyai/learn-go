#!/usr/bin/env bash
rm -f private/ca.key.pem
openssl genrsa -aes256 -out private/ca.key.pem 4096
chmod 400 private/ca.key.pem
#!/usr/bin/env bash

HOST=api.stefarf.com

rm -f intermediate/private/$HOST.key.pem
openssl genrsa -out intermediate/private/$HOST.key.pem 2048 # without password
#openssl genrsa -aes256 -out intermediate/private/$HOST.key.pem 2048 # with password
chmod 400 intermediate/private/$HOST.key.pem
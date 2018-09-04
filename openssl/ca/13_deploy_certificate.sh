#!/usr/bin/env bash
HOST=api.stefarf.com
cp \
intermediate/certs/ca-chain.cert.pem \
intermediate/private/$HOST.key.pem \
intermediate/certs/$HOST.cert.pem \
..
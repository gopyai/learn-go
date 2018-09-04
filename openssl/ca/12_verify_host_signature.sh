#!/usr/bin/env bash
HOST=api.stefarf.com
openssl verify -CAfile intermediate/certs/ca-chain.cert.pem intermediate/certs/$HOST.cert.pem
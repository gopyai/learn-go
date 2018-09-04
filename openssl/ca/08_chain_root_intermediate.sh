#!/usr/bin/env bash
rm -f intermediate/certs/ca-chain.cert.pem
cat intermediate/certs/intermediate.cert.pem certs/ca.cert.pem > intermediate/certs/ca-chain.cert.pem
chmod 444 intermediate/certs/ca-chain.cert.pem
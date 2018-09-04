#!/usr/bin/env bash
openssl verify -CAfile certs/ca.cert.pem intermediate/certs/intermediate.cert.pem

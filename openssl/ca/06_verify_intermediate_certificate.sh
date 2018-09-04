#!/usr/bin/env bash
openssl x509 -noout -text -in intermediate/certs/intermediate.cert.pem > verify_intermediate.txt
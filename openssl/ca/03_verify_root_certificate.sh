#!/usr/bin/env bash
openssl x509 -noout -text -in certs/ca.cert.pem > verify_root.txt
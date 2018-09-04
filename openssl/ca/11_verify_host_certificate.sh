#!/usr/bin/env bash
HOST=api.stefarf.com
openssl x509 -noout -text -in intermediate/certs/$HOST.cert.pem > verify_host.txt
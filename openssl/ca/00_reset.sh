#!/usr/bin/env bash

rm -f \
index.txt \
index.txt.attr \
index.txt.old \
serial \
serial.old \
certs/* \
crl/* \
newcerts/* \
private/* \
intermediate/certs/* \
intermediate/crl/* \
intermediate/csr/* \
intermediate/newcerts/* \
intermediate/private/* \
verify_*.txt

touch index.txt
touch serial
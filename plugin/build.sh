#!/bin/sh

(cd plug && go build -buildmode=plugin -o ../plug.so)
(cd main  && go build -o ../run)
./run
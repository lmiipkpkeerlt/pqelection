#!/usr/bin/env bash

if [ $# -eq 0 ]; then
    echo "No argument supplied"
    exit 1
fi

client=$1

# Generate client’s private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes \
    -keyout ${client}-key.pem \
    -out ${client}-req.pem \
    -subj "/C=/ST=/L=/O=/OU=/CN=/emailAddress="

# Sign the Client Certificate Request (CSR)
openssl x509 -req -in ${client}-req.pem -days 60 \
    -CA ca-cert.pem \
    -CAkey ca-key.pem \
    -CAcreateserial \
    -out ${client}-cert.pem

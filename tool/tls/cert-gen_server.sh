#!/usr/bin/env bash

if [ $# -eq 0 ]; then
    echo "No argument supplied"
    exit 1
fi

server=$1

# Generate Web Server’s Private Key and CSR (Certificate Signing Request)
# -x509 removed because we don't want to self-sign certificate as like as CA
# certificate.
# -days removed deleted because we are creating CSR instead of certificate
# -keyout name of output key
# -out name of certificate request
openssl req -newkey rsa:4096 -nodes \
    -keyout ${server}-key.pem \
    -out ${server}-req.pem \
    -subj "/C=/ST=/L=/O=/OU=/CN=localhost/emailAddress="

# Sign the Web Server Certificate Request (CSR)
# -req we are going to pass in certificate request
# -inflag the name of the request file that is server-req.pem
# -CAflag pass Certificate File of CA:ca-cert.pem
# -CAKey pass private key of CA: ca-key.pem
# -CAcreateserial CA must ensure that each certificate it signs with a unique
# serial number
# -out output to certificate file: server-cert.pem
openssl x509 -req \
    -in ${server}-req.pem \
    -CA ca-cert.pem \
    -CAkey ca-key.pem \
    -CAcreateserial \
    -out ${server}-cert.pem
    #-extfile <(printf "subjectAltName=DNS:localhost") \

#!/usr/bin/env bash

# Generate CA private key and self-signed certificate 
# -x509 provides self-signed certificate instead of a certificate request
# -newkey rsa:4096 provides both RSA key with 4096-bit and its certificate
# request at the same timeA
# -nodes to not encrypt the private key
# -days certificate valid date.
# -keyout write the created private key to ca-key.pem file
# -out write the certificate to ca-cert.pem file
# -subj:
# /C=Country
# /ST=State or province
# /L=Locality name or city
# /O=Organization
# /OU=Organization Unit
# /CN=Common Name or domain name
# /emailAddress=...
openssl req -x509 -newkey rsa:4096 -nodes -days 365 \
    -keyout ca-key.pem \
    -out ca-cert.pem \
    -subj "/C=/ST=/L=/O=/OU=/CN=/emailAddress="

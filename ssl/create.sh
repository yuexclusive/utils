#!/bin/bash

# Create the server CA certs.
openssl req -x509                                     \
  -newkey rsa:4096                                    \
  -nodes                                              \
  -days 3650                                          \
  -keyout server_ca_key.pem                           \
  -out server_ca_cert.pem                             \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server_ca/   \
  -config ./openssl.cnf                               \
  -extensions test_ca

# Create the client CA certs.
openssl req -x509                                     \
  -newkey rsa:4096                                    \
  -nodes                                              \
  -days 3650                                          \
  -keyout client_ca_key.pem                           \
  -out client_ca_cert.pem                             \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client_ca/   \
  -config ./openssl.cnf                               \
  -extensions test_ca

# Generate server certs.
openssl genrsa -out server_key.pem 4096
openssl req -new                                    \
  -key server_key.pem                              \
  -days 3650                                        \
  -out server_csr.pem                              \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server/   \
  -config ./openssl.cnf                             \
  -reqexts test_server
openssl x509 -req           \
  -in server_csr.pem       \
  -CAkey server_ca_key.pem  \
  -CA server_ca_cert.pem    \
  -days 3650                \
  -set_serial 1000          \
  -out server_cert.pem     \
  -extfile ./openssl.cnf    \
  -extensions test_server
openssl verify -verbose -CAfile server_ca_cert.pem  server_cert.pem

# Generate two client certs.
openssl genrsa -out client_key.pem 4096
openssl req -new                                    \
  -key client_key.pem                              \
  -days 3650                                        \
  -out client_csr.pem                              \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client/   \
  -config ./openssl.cnf                             \
  -reqexts test_client
openssl x509 -req           \
  -in client_csr.pem       \
  -CAkey client_ca_key.pem  \
  -CA client_ca_cert.pem    \
  -days 3650                \
  -set_serial 1000          \
  -out client_cert.pem     \
  -extfile ./openssl.cnf    \
  -extensions test_client
openssl verify -verbose -CAfile client_ca_cert.pem  client_cert.pem

# Generate a cert with SPIFFE ID.
openssl req -x509                                                         \
  -newkey rsa:4096                                                        \
  -keyout spiffe_key.pem                                                  \
  -out spiffe_cert.pem                                                    \
  -nodes                                                                  \
  -days 3650                                                              \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client/                         \
  -addext "subjectAltName = URI:spiffe://foo.bar.com/client/workload/1"

# Generate a cert with SPIFFE ID and another SAN URI field(which doesn't meet SPIFFE specs).
openssl req -x509                                                         \
  -newkey rsa:4096                                                        \
  -keyout multiple_uri_key.pem                                            \
  -out multiple_uri_cert.pem                                              \
  -nodes                                                                  \
  -days 3650                                                              \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client/                         \
  -addext "subjectAltName = URI:spiffe://foo.bar.com/client/workload/1, URI:https://bar.baz.com/client"
# Cleanup the CSRs.
rm *_csr.pem

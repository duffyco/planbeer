#!/bin/sh

cat > $3/req.cnf <<EOF
[v3_req]
keyUsage = critical, digitalSignature, keyAgreement
extendedKeyUsage = serverAuth
subjectAltName = @alt_names
[alt_names]
DNS.1 = $1
DNS.2 = $2
EOF

openssl req -x509 -sha256 -newkey rsa:2048 -nodes -keyout $3/domain.key -days 1825  -out  $3/domain.crt  -subj "/CN=Plan B-eer Root CA"

openssl req -newkey rsa:2048 -nodes -subj "/CN=picobrew.com" \
      -keyout  $3/server.key -out  $3/server.csr

openssl x509 \
        -CA $3/domain.crt -CAkey $3/domain.key -CAcreateserial \
       -in $3/server.csr \
       -req -days 1825 -out  $3/server.crt  -extfile $3/req.cnf -extensions v3_req

cat  $3/server.crt  $3/domain.crt >  $3/bundle.crt


#!/usr/bin/env bash
set -e

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
CUR_DIR=$(pwd)

cd $BASE_DIR/certificates

# openssl req -x509 \
#   -newkey rsa:4096 \
#   -keyout key.pem \
#   -out cert.pem \
#   -days 365 \
#   -subj "/C=US/ST=California/L=San Diego/O=Versions Club/OU=Members/CN=versions.local"
which certstrap

echo "init cert authority"
certstrap init --common-name "Versions Local CA" \
  --expires "2 years" \
  --organization "Versions Local" \
  --organizational-unit "Creators" \
  --country "US" \
  --province "California" \
  --locality "San Diego" \
  --passphrase ""
  # --depot-path "$BASE_DIR/certificates"

echo "request certs"
certstrap request-cert --common-name "versions.local" \
  --ip 127.0.0.1 \
  --domain *.versions.local,versions.local \
  --organization "Versions Local" \
  --organizational-unit "Creators" \
  --country "US" \
  --province "California" \
  --locality "San Diego" \
  --passphrase ""

echo "sign certs"
certstrap sign "versions.local" \
  --CA "Versions Local CA" \
  --expires "2 years" \
  --passphrase ""
  # --intermediate 




cd $CUR_DIR
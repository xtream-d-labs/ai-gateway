#!/bin/bash

if [ -z "${AIG_JWS_PRIVATE_KEY}" ]; then
  AIG_JWS_PRIVATE_KEY=/certs/private.pem
  export AIG_JWS_PRIVATE_KEY
fi
if [ ! -e "${AIG_JWS_PRIVATE_KEY}" ]; then
  mkdir -p "$( dirname "${AIG_JWS_PRIVATE_KEY}" )"
  openssl genrsa -out "${AIG_JWS_PRIVATE_KEY}"
fi
if [ -z "${AIG_JWS_PUBLIC_KEY}" ]; then
  AIG_JWS_PUBLIC_KEY=/certs/public.pem
  export AIG_JWS_PUBLIC_KEY
fi
if [ ! -e "${AIG_JWS_PUBLIC_KEY}" ]; then
  mkdir -p "$( dirname "${AIG_JWS_PUBLIC_KEY}" )"
  openssl rsa -in "${AIG_JWS_PRIVATE_KEY}" -pubout -out "${AIG_JWS_PUBLIC_KEY}"
fi

exec /app --scheme http --host 0.0.0.0 --port 80

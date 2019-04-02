#!/bin/bash

if [ -z "${SS_JWS_PRIVATE_KEY}" ]; then
  SS_JWS_PRIVATE_KEY=/certs/private.pem
  export SS_JWS_PRIVATE_KEY
fi
if [ ! -e "${SS_JWS_PRIVATE_KEY}" ]; then
  mkdir -p $( dirname "${SS_JWS_PRIVATE_KEY}" )
  openssl genrsa -out "${SS_JWS_PRIVATE_KEY}"
fi
if [ -z "${SS_JWS_PUBLIC_KEY}" ]; then
  SS_JWS_PUBLIC_KEY=/certs/public.pem
  export SS_JWS_PUBLIC_KEY
fi
if [ ! -e "${SS_JWS_PUBLIC_KEY}" ]; then
  mkdir -p $( dirname "${SS_JWS_PUBLIC_KEY}" )
  openssl rsa -in "${SS_JWS_PRIVATE_KEY}" -pubout -out "${SS_JWS_PUBLIC_KEY}"
fi

exec /app --scheme http --host 0.0.0.0 --port 80

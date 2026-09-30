#!/usr/bin/env bash
set -euo pipefail

: "${RAPIDAPI_KEY:?Set RAPIDAPI_KEY before running this example}"

host="florida-sunbiz-entity-and-officer-lookup.p.rapidapi.com"

curl --fail-with-body --silent --show-error \
  --get "https://${host}/v1/entity/search" \
  --data-urlencode "name=publix" \
  --data-urlencode "limit=2" \
  --header "X-RapidAPI-Key: ${RAPIDAPI_KEY}" \
  --header "X-RapidAPI-Host: ${host}"

printf '\n'

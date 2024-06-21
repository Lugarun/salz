#!/usr/bin/env bash

SALZ_SECRET_LOCATION="../../salz-secrets"
SALZ_SECRETS=$($SALZ_SECRET_LOCATION/decrypt.sh $SALZ_SECRET_LOCATION/dev-secrets.json.age)

export ENV="local"

export OAUTH_CLIENT_KEY=$(echo $SALZ_SECRETS | jq -r '.oauth_client_id')
export OAUTH_CLIENT_SECRET=$(echo $SALZ_SECRETS | jq -r '.oauth_client_secret')

# Local environment variables
export PG_URI="postgres://salz:superdupersecret@localhost:5432/salz?sslmode=disable"
 
go run .
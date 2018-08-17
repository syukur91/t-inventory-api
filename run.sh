#!/bin/bash

echo "building t-inventory-api ..."

if CGO_ENABLED=0 go build -o ./bin/t-inventory-api
    then
        echo "completed building t-inventory-api"

      ./bin/t-inventory-api
    else
        echo "failed to build t-inventory-api"
fi
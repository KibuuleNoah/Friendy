#!/bin/bash





AUTH_HEAD="Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2xsZWN0aW9uSWQiOiJfcGJfdXNlcnNfYXV0aF8iLCJleHAiOjE3NjI4NzE0ODEsImlkIjoiMTEwZjh6ODBwaTZzMmIwIiwicmVmcmVzaGFibGUiOnRydWUsInR5cGUiOiJhdXRoIn0.W2AnGlqqoZmvHmifohdfDqkipTW3iEpr3XXkqADee98"
BASE_URL="http://localhost:8090/api/friendy/"

if [ "$1" == "a" ]; then
  # Get All user friends
  curl -X GET ${BASE_URL}v1/all/ \
    -H "$AUTH_HEAD"
    -H "Content-Type: application/json" \

elif [ "$1" == "c" ]; then
  # Create Friend
  curl -X POST ${BASE_URL}v1/friend/ \
    -H "$AUTH_HEAD"\
    -H "Content-Type: application/json" \
    -d '{"fullname":"Tristar Moxie","tel":"+25671346778","desc":"A friend from last work","first_met_on":"2025-11-04","met_place":"Masaka","tags":"work, friend"}'

elif [ "$1" == "r" ]; then
  # Get One Friend
  curl -X GET ${BASE_URL}v1/friend/ \
    -H "$AUTH_HEAD"\
    -H "Content-Type: application/json" \
    -d '{"id":"7wm175x4ry5znvq"}'

elif [ "$1" == "u" ]; then
  # Update Friend
  curl -X PUT ${BASE_URL}v1/friend/ \
    -H "$AUTH_HEAD"\
    -H "Content-Type: application/json" \
    -d '{"id":"7wm175x4ry5znvq","fullname":"Star Nahum Kibuule","desc":"This is Code bbro forever","created":"xxxxxxd"}'
  #,"tel":"+25671346778","desc":"A friend from last work","first_met_on":"2025-11-04","met_place":"Masaka","tags":"work, friend"}'

elif [ "$1" == "d" ]; then
  # Update Friend
  curl -X DELETE ${BASE_URL}v1/friend/ \
    -H "$AUTH_HEAD"\
    -H "Content-Type: application/json" \
    -d '{"id":"1b3fzbkzs1itrt5"}'

else
    echo "Unknown argument: $1"
fi

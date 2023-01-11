#!/bin/bash

set -o xtrace
set -o nounset
set -o errexit
set -o pipefail

trap "trap - SIGTERM && kill -- -$$" SIGINT SIGTERM EXIT

RUN_SERVER=(
    "go"
    "run"
    "./server"
)

RUN_CLIENT=(
    "go"
    "run"
    "./client"
)

${RUN_SERVER[@]} --grpcport 5001 --servername server1 &
${RUN_SERVER[@]} --grpcport 5002 --servername server2 &
${RUN_SERVER[@]} --grpcport 5003 --servername server3 &

sleep 3

time ${RUN_CLIENT[@]} -n 100 \
    -server localhost:5001 \
    -server localhost:5002 \
    -server localhost:5003

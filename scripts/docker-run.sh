#!/bin/bash

set -e
set -u

sudo docker run -p '8001:8001' \
    -e 'ENV_NAME=test' \
    -e 'LOG_LEVEL=warn' \
    -e 'MAX_RANDOM_DELAY_MS=1000' \
    fake-service

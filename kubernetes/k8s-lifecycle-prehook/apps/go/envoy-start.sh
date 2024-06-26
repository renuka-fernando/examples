#!/bin/bash
set -e

echo "Starting Envoy..."

_term() {
    echo "Stopping Envoy..."
    kill -SIGTERM $ENVOY_PID
    wait $ENVOY_PID
    echo "Envoy stopped."
    exit 0
}

/usr/local/bin/envoy \
    -c /etc/envoy/envoy.yaml &

ENVOY_PID=$!
# trap handle_signal SIGTERM SIGINT
trap _term SIGTERM SIGINT

wait $ENVOY_PID

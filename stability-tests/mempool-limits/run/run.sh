#!/bin/bash

APPDIR=/tmp/anumad-temp
ANUMAD_RPC_PORT=29587

rm -rf "${APPDIR}"

anumad --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${ANUMAD_RPC_PORT}" --profile=6061 &
ANUMAD_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${ANUMAD_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $ANUMAD_PID

wait $ANUMAD_PID
ANUMAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Anumad exit code: $ANUMAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ANUMAD_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1

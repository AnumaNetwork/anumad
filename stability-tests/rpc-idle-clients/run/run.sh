#!/bin/bash
rm -rf /tmp/anumad-temp

NUM_CLIENTS=128
anumad --devnet --appdir=/tmp/anumad-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
ANUMAD_PID=$!
ANUMAD_KILLED=0
function killAnumadIfNotKilled() {
  if [ $ANUMAD_KILLED -eq 0 ]; then
    kill $ANUMAD_PID
  fi
}
trap "killAnumadIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $ANUMAD_PID

wait $ANUMAD_PID
ANUMAD_EXIT_CODE=$?
ANUMAD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Anumad exit code: $ANUMAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ANUMAD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1

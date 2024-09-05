#!/bin/bash
rm -rf /tmp/anumad-temp

anumad --devnet --appdir=/tmp/anumad-temp --profile=6061 --loglevel=debug &
ANUMAD_PID=$!
ANUMAD_KILLED=0
function killAnumadIfNotKilled() {
    if [ $ANUMAD_KILLED -eq 0 ]; then
      kill $ANUMAD_PID
    fi
}
trap "killAnumadIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $ANUMAD_PID

wait $ANUMAD_PID
ANUMAD_KILLED=1
ANUMAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Anumad exit code: $ANUMAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ANUMAD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1

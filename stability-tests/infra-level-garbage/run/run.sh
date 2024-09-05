#!/bin/bash
rm -rf /tmp/anumad-temp

anumad --devnet --appdir=/tmp/anumad-temp --profile=6061 &
ANUMAD_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:16611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $ANUMAD_PID

wait $ANUMAD_PID
ANUMAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Anumad exit code: $ANUMAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ANUMAD_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1

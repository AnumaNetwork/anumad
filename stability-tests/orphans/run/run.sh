#!/bin/bash
rm -rf /tmp/anumad-temp

anumad --simnet --appdir=/tmp/anumad-temp --profile=6061 &
ANUMAD_PID=$!

sleep 1

orphans --simnet -alocalhost:16511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $ANUMAD_PID

wait $ANUMAD_PID
ANUMAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Anumad exit code: $ANUMAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ANUMAD_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1

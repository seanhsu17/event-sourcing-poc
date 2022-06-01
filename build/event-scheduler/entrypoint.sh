#!/bin/sh

cmd="${WORK_DIR}/main"
echo -----------------------------------------------------
date
$cmd &
child=$!
wait $child
wait $child
exit_status=$?
echo end of child process with code $exit_status

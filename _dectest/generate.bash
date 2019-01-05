#!/usr/bin/env bash

set -euo pipefail

if [ "${TRAVIS_EVENT_TYPE}" == "cron" ]; then
	pushd _dectest
	curl -sS http://speleotrove.com/decimal/dectest.zip > dectest.zip
  unzip dectest.zip
  rm dectest.zip
	popd
fi

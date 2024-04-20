#!/bin/sh

cd proto
buf lint || exit $?

buf generate --template buf.gen.yaml

cd ../pkg

buf format -w pb


#!/bin/sh

call_buf_generate() {
  buf generate "$@" || exit $?
}

# 開始時刻を保存
touch -t $(date +%Y%m%d%H%M.%S) /tmp/buf-start-timestamp

##
# start proto/source
##
cd proto
buf lint || exit $?

# api
call_buf_generate --template buf.gen.api.yaml \
  --path proto/rpc/api


cd - > /dev/null
##
# end proto/source
##

# 古いファイルの削除
find proto/generated/ -name '*.gen.proto' ! -newer /tmp/buf-start-timestamp -exec rm {} \;

# 古いファイルの削除
find pkg/ \( -name '*.pb.go' -o -name '*.pb.validate.go' -o -name '*.connect.go' -o -name '*.gen.go' \) ! -newer /tmp/buf-start-timestamp -exec rm {} \;

buf format -w pb

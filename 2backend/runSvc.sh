#!/bin/bash

trap "exit" INT TERM ERR
trap "kill 0" EXIT

#svcs
#go run svc/ts/cmd/ts-server/main.go --host 0.0.0.0 --port 8002 --scheme http &

#gateway
#go run svc/gateway/main.go &
go run svc/gateway/lugat/main.go &

#consumer
go run mware/consumer/main.go &

wait
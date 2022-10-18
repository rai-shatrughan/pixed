#!/bin/bash

trap "exit" INT TERM ERR
trap "kill 0" EXIT

#timeseries
go run svc/timeseries/cmd/main.go &

#gateway
#go run svc/gateway/cmd/main.go &
go run svc/gateway/lugat/main.go &

#consumer
go run pkg/consumer/main.go &

wait
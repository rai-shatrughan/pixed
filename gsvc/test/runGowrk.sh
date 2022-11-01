#!/bin/bash

#GSVC
HOST=172.18.0.21:8000
#RSVC
#HOST=172.18.0.22:9000

putTS(){
    echo ""
    echo "---put TS---"
    echo ""

    start=`date`
    echo "Start Time - " $start

    go-wrk -c 2000 -d 3600 -T 30000 -M POST \
        -H "X-API-Key: sr12345" \
        -H "Content-Type: application/json" \
        -body @test/json/ex.json \
        http://${HOST}/api/mindconnect/v3/exchange/6fdae6af-226d-48bd-8b61-699758137eb3

    end=`date`
    echo "End Time - " $end
}

getTS(){
    echo ""
    echo "***get TS***"
    echo ""

    start=`date`
    echo "Start Time - " $start

        go-wrk -c 500 -d 30 -T 50000 -M GET \
        -H "X-API-Key: sr12345" \
        http://172.18.0.21:8000/api/iottimeseries/v3/timeseries/6fdae6af-226d-48bd-8b61-699758137eb3?duration=1m

    end=`date`
    echo "End Time - " $end
}

getCurlTS(){
    curl -H "X-API-Key: sr12345" \
    http://172.18.0.21:8000/api/iottimeseries/v3/timeseries/6fdae6af-226d-48bd-8b61-699758137eb3?duration=1m
}

putTS
#getTS

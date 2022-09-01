#!/bin/bash
start=`date`
echo "Start Time - " $start

go-wrk -c 4000 -d 60 -T 30000 -M PUT \
    -H "Authorization: Bearer sr12345" \
    -H "Content-Type: application/json" \
    -body @json/ts.json \
    http://172.18.0.21:8002/api/iottimeseries/v3/timeseries

end=`date`
echo "End Time - " $end

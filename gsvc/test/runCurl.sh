exCurlTS(){
    curl -H "X-API-Key: sr12345" \
    -X POST \
    -d @test/json/ts.json \
    http://172.18.0.21:8000/api/v1/exchange/6fdae6af-226d-48bd-8b61-699758137eb3
}

exCurlTS
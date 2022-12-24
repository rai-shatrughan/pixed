mkdir api
cd api

mkdir api/agm
cd api/agm
swagger generate server -f ../../yaml/agentmanagement-v3-4-0.yaml -A agm -P models.Principal
cd ../../
go get -u -f ./...
go run api/agm/cmd/agm-server/main.go --host 0.0.0.0 --port 8001 --scheme http

mkdir api/ts
cd api/ts
swagger generate server -f ../../yaml/iottimeseries-v3-7-0.yaml -A ts -P models.Principal
cd ../../
go mod tidy
go run api/ts/cmd/ts-server/main.go --host 0.0.0.0 --port 8002 --scheme http

mkdir api/am
cd api/am
swagger generate server -f ../../yaml/assetmanagement-v3-18-3.yaml -A am -P models.Principal
cd ../../
go mod tidy
go run api/am/cmd/am-server/main.go --host 0.0.0.0 --port 8003 --scheme http
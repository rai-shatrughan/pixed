.PHONY: clean test perf bench

clean:
	rm -rf ../docker/app/bin
	rm -rf ../docker/app/conf

test:
	go test test/exch_test.go -args -host gsvc

perf:
	bash test/runGowrk.sh

bench:
	go test test/exch_test.go -bench=. -benchtime=10s -args -host gsvc

gsvc:
	cp -r pkg/conf ../docker/app
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ../docker/app/bin/gsvc server/main.go
	docker-compose -f ../docker/docker-compose.yml --env-file ../docker/.env build gsvc 
	docker-compose -f ../docker/docker-compose.yml --env-file ../docker/.env up -d gsvc

cons:
	cp -r pkg/conf ../docker/app
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ../docker/app/bin/consumer pkg/consumer/main.go
	docker-compose -f ../docker/docker-compose.yml --env-file ../docker/.env build cons 
	docker-compose -f ../docker/docker-compose.yml --env-file ../docker/.env up -d cons

all:
	bash ../docker/buildNstart.sh
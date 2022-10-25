.PHONY: clean test perf bench

clean:
	rm -rf docker/app/bin
	rm -rf docker/app/conf

test:
	cd gsvc; \
	go test test/exch_test.go -args -host gsvc

testr:
	cd gsvc; \
	go test test/exch_test.go -args -host rsvc

perf:
	cd gsvc; \
	bash test/runGowrk.sh

bench:
	cd gsvc; \
	go test test/exch_test.go -bench=. -benchtime=10s -args -host gsvc

gsvc:
	cd gsvc; \
	cp -r pkg/conf ../docker/app; \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ../docker/app/bin/gsvc server/main.go
	docker-compose -f docker/docker-compose.yml --env-file docker/.env build gsvc 
	docker-compose -f docker/docker-compose.yml --env-file docker/.env up -d gsvc

rsvc:
	cd rsvc; \
	cargo install --target x86_64-unknown-linux-musl --path . ; \
	cp target/x86_64-unknown-linux-musl/release/rsvc ../docker/app/bin/rsvc
	docker-compose -f docker/docker-compose.yml --env-file docker/.env build rsvc 
	docker-compose -f docker/docker-compose.yml --env-file docker/.env up -d rsvc

cons:
	cd gsvc; \
	cp -r pkg/conf ../docker/app; \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ../docker/app/bin/consumer pkg/consumer/main.go
	docker-compose -f docker/docker-compose.yml --env-file docker/.env build cons 
	docker-compose -f docker/docker-compose.yml --env-file docker/.env up -d cons

all:
	cd docker; \
	bash buildNstart.sh
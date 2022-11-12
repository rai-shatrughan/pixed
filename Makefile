.PHONY: clean test perf bench gsvc rsvc cons

all: cnd down dirs appbin buildd up clean

#docker start

cnd:	
	docker network inspect sr_cluster_network >/dev/null 2>&1 || \
    docker network create --driver=bridge --subnet=172.18.0.0/23 sr_cluster_network

down:
	cd docker; \
	docker-compose --env-file .env -f docker-compose.yml down

dirs:
	base=/data; \
	sudo rm -rf $$base; \
    for dir in sql kafka kafka_zk hbase hbase_zk prometheus grafana; do \
		wd=$$base/$$dir; \
        echo $$wd; \
		sudo mkdir -p $$wd; \
		sudo chown nobody:nogroup $$wd; \
		sudo chmod 777 $$wd -R; \
		echo "Created " $$wd; \
    done

appbin:
	cd gsvc; \
	cp -r pkg/conf ../docker/app; \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ../docker/app/bin/consumer pkg/consumer/main.go; \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ../docker/app/bin/gsvc server/main.go
	cd rsvc; \
	cargo install --path . ; \
	cp target/release/rsvc ../docker/app/bin/rsvc

buildd:
	cd docker; \
	docker-compose --env-file .env -f docker-compose.yml build

up:
	cd docker; \
	docker-compose --env-file .env -f docker-compose.yml up -d

rm:
	docker stop $$(docker ps -aq)
	docker rm $$(docker ps -aq)

clean:
	rm -rf docker/app/bin
	rm -rf docker/app/conf
	docker image prune -f

#docker end

testv:
	cd gsvc; \
	go test -v test/ingest_test.go -args -host gsvc

test:
	cd gsvc; \
	go test test/ingest_test.go -args -host gsvc

bench:
	cd gsvc/test; \
	go test -bench=. -v

perfn:
	cd gsvc; \
	go run test/cmd_new/perf_new.go -host gsvc

perfo:
	cd gsvc; \
	go run test/cmd_old/perf_old.go -host gsvc

wrk:
	cd gsvc; \
	bash test/runGowrk.sh

gsvc:
	cd gsvc; \
	cp -r pkg/conf ../docker/app; \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ../docker/app/bin/gsvc server/main.go
	docker-compose -f docker/docker-compose.yml --env-file docker/.env build gsvc 
	docker-compose -f docker/docker-compose.yml --env-file docker/.env up -d gsvc
	$(MAKE) clean

rsvc:
	cd rsvc; \
	mkdir -p ../docker/app/bin/; \
	PKG_CONFIG_ALLOW_CROSS=true PKG_CONFIG_ALL_STATIC=true LIBZ_SYS_STATIC=1 CC=musl-gcc CXX=g++ cargo build --release --target x86_64-unknown-linux-musl
	cp rsvc/target/x86_64-unknown-linux-musl/release/rsvc docker/app/bin/rsvc
	docker-compose -f docker/docker-compose.yml --env-file docker/.env build rsvc 
	docker-compose -f docker/docker-compose.yml --env-file docker/.env up -d rsvc
	$(MAKE) clean

cons:
	cd gsvc; \
	cp -r pkg/conf ../docker/app; \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ../docker/app/bin/consumer pkg/consumer/main.go
	docker-compose -f docker/docker-compose.yml --env-file docker/.env build cons 
	docker-compose -f docker/docker-compose.yml --env-file docker/.env up -d cons
	$(MAKE) clean



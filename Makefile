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
    for dir in kafka zookeeper sql etcd1 prometheus grafana; do \
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
	cargo install --target x86_64-unknown-linux-musl --path . ; \
	cp target/x86_64-unknown-linux-musl/release/rsvc ../docker/app/bin/rsvc

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
	$(MAKE) clean

rsvc:
	cd rsvc; \
	cargo install --target x86_64-unknown-linux-musl --path . ; \
	cp target/x86_64-unknown-linux-musl/release/rsvc ../docker/app/bin/rsvc
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



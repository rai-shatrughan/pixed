 #!/bin/bash
 
root_path=/data
data_folders=("kafka" "zookeeper" "sql" "etcd1")

1create_network() {
    docker network create \
        --driver=bridge \
        --subnet=172.18.0.0/23 \
        sr_cluster_network
}

2down_containers() {
    docker-compose --env-file .env -f docker-compose.yml down
}

remove_dir() {
    wd=$1    
    if [[ -d $wd ]]; then
        echo "Removing " $wd
        sudo rm -rf $wd
    fi
}

crate_dir(){
    wd=$1
    sudo mkdir -p $wd
    sudo chown nobody:nogroup $wd
    sudo chmod 777 $wd -R
    echo "Created " $wd
}

3recreate_dir(){
    for dir in ${data_folders[@]}; do
        wd=$root_path/$dir    
        remove_dir $wd
        crate_dir  $wd  
    done
}

4build_go_svc(){
    cd ../gsvc
    echo "building go binary"

    cp -r pkg/conf ../docker/app
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ../docker/app/bin/consumer pkg/consumer/main.go
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ../docker/app/bin/gsvc server/main.go
    cd ../docker
}

4build_rust_svc(){
    cd ../rsvc
    echo "building rust binary"

    cargo install --target x86_64-unknown-linux-musl --path .
       
    cp target/x86_64-unknown-linux-musl/release/rsvc ../docker/app/bin/rsvc 
    
    cd ../docker
}

5recreate_containers() {
    docker-compose --env-file .env -f docker-compose.yml build 
    docker-compose --env-file .env -f docker-compose.yml up -d
}

6setup_db() {
    echo "create Topic*****"
    echo
    bash kafka/createTopic.sh
}

7clean_app_bin(){
    echo "removing app/bin/ and app/conf"
    rm -rf app/bin
    rm -rf app/conf
}

1create_network
2down_containers
3recreate_dir
4build_go_svc
4build_rust_svc
5recreate_containers
#6setup_db
7clean_app_bin

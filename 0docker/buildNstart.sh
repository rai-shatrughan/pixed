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

4build_app_bin(){
    cd ../2backend
    echo "building go binary"

    cp -r mware/conf ../0docker/app
    go build -ldflags "-linkmode external -extldflags -static" -o ../0docker/app/bin/consumer -a mware/consumer/main.go
    go build -ldflags "-linkmode external -extldflags -static" -o ../0docker/app/bin/ts-server -a svc/ts/cmd/ts-server/main.go
    #go build -o ../0docker/app/bin/consumer mware/consumer/main.go
    #go build -o ../0docker/app/bin/ts-server svc/ts/cmd/ts-server/main.go 
    cd ../0docker
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
4build_app_bin
5recreate_containers
#6setup_db
7clean_app_bin
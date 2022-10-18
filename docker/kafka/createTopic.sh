#!/bin/bash
replicationFactor=1 #no. of brokers
partition=5
topics=("websvc" "ts")


partition_bigdata=5
topic_bigdata=("tsbig")

function createTopics() {
    for topic in ${topics[@]};
    do
    echo "Creating Topic : " $topic

    docker exec -t kafka /bin/sh -c \
    "/opt/bitnami/kafka/bin/kafka-topics.sh --create --bootstrap-server 172.18.0.41:9092 \
    --replication-factor $replicationFactor --partitions $partition \
    --topic $topic"

    echo "Topic Created : " $topic " - Partition Count : " $partition
    done
}

function createBigTopics() {
    for topic in ${topic_bigdata[@]};
    do
    echo "Creating Topic : " $topic

    docker exec -t kafka /bin/sh -c \
    "/opt/bitnami/kafka/bin/kafka-topics.sh --create --bootstrap-server 172.18.0.41:9092 \
    --replication-factor $replicationFactor  --partitions $partition_bigdata \
    --topic $topic"

    echo "Topic Created : " $topic " - Partition Count : " $partition_bigdata
    done
}

createTopics
# createBigTopics


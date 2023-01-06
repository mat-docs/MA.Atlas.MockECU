# Usage:
# set environment variable called KAFKA_CREATE_TOPICS to:
# <topic_name>:<num_partitions>[:<normal|compact>],<topic_name>:...
# log compaction defaults to "normal" (not compacted).
IFS=","; for topic in $KAFKA_CREATE_TOPICS; do
    echo -n "creating topic: $topic"
    IFS=':' read -r -a config <<< "$topic"

    if [ "${config[2]}" == "compact" ]; then
        echo " (WITH log compaction)"
        /usr/bin/kafka-topics --create \
            --zookeeper $KAFKA_ZOOKEEPER_CONNECT \
            --replication-factor ${KAFKA_REPLICATION_FACTOR-3} --partitions ${config[1]} \
            --config "cleanup.policy=compact" \
            --config "delete.retention.ms=86400000" \
            --config "max.message.bytes=2097164" \
            --config "retention.bytes=-1" \
            --config "retention.ms=2419200000" \
            --topic ${config[0]}
    else
        echo " (no log compaction)"
        /usr/bin/kafka-topics --create \
            --zookeeper $KAFKA_ZOOKEEPER_CONNECT \
            --replication-factor ${KAFKA_REPLICATION_FACTOR-3} --partitions ${config[1]} \
            --topic ${config[0]}
    fi
done

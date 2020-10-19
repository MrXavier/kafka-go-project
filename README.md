# kafka-go-project
Project to implement kafka integration in Golang

### How run
#### Start Zookepper server
zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties
#### Start Kafka server
kafka-server-start /usr/local/etc/kafka/server.properties

#### Create a topic
kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic TOPIC_NAME
#### List topic
kafka-topics --list --bootstrap-server localhost:9092
#### Describe topic
kafka-topics --describe --bootstrap-server localhost:9092 --topic TOPIC

### Read from topic
kafka-console-consumer --bootstrap-server localhost:9092 --topic TOPIC --from-beginning

### Write to topic
kafka-console-producer --bootstrap-server localhost:9092 --topic TOPIC

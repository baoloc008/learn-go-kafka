# Learn Go Kafka

## Links
- [Lập trình Golang kết nối Kafka](https://techmaster.vn/posts/34654/lap-trinh-golang-ket-noi-kafka)
- [Part 1: Apache Kafka for beginners - What is Apache Kafka? - CloudKarafka, Apache Kafka Message streaming as a Service](https://www.cloudkarafka.com/blog/2016-11-30-part1-kafka-for-beginners-what-is-apache-kafka.html)
- [Getting Started with Kafka in Golang - Yusuf Syaifudin - Medium](https://medium.com/@yusufs/getting-started-with-kafka-in-golang-14ccab5fa26)

## Run
#### Run zookeeper and kafka
```sh
MY_IP=<IP> docker-compose up 
```
#### Create topic
```sh
docker run --net=host --rm confluentinc/cp-kafka kafka-topics --create --topic LocTopic --partitions 3 --replication-factor 2 --if-not-exists --zookeeper localhost:32181
```
#### Listen "LocTopic" in partition 2
```sh
kafkacat -C -b localhost:19092,localhost:29092,localhost:39092 -t LocTopic -p 2
```
#### Publish message to "LocTopic" partition 2
```sh
echo 'Loc call api' | kafkacat -P -b localhost:19092,localhost:29092,localhost:39092 -t LocTopic -p 2
```

# Learn Go Kafka

## Links
- [Lập trình Golang kết nối Kafka](https://techmaster.vn/posts/34654/lap-trinh-golang-ket-noi-kafka)
- [Part 1: Apache Kafka for beginners - What is Apache Kafka? - CloudKarafka, Apache Kafka Message streaming as a Service](https://www.cloudkarafka.com/blog/2016-11-30-part1-kafka-for-beginners-what-is-apache-kafka.html)
- [Getting Started with Kafka in Golang - Yusuf Syaifudin - Medium](https://medium.com/@yusufs/getting-started-with-kafka-in-golang-14ccab5fa26)
- [Understanding Kafka Topics and Partitions - Stack Overflow](https://stackoverflow.com/questions/38024514/understanding-kafka-topics-and-partitions)
- [What is Apache Kafka \| CodeFlex](http://codeflex.co/what-is-apache-kafka/)

## Run
#### Run zookeeper and kafka
```sh
MY_IP=<IP> docker-compose up 
```

#### Create topic
```sh
docker run --net=host --rm confluentinc/cp-kafka kafka-topics --create --topic DemoTopic --partitions 3 --replication-factor 2 --if-not-exists --zookeeper localhost:32181
```

#### Delete topic
```sh
docker run --net=host --rm confluentinc/cp-kafka kafka-topics --delete --topic DemoTopic --zookeeper localhost:32181
```

#### Listen "DemoTopic" in partition 2
```sh
kafkacat -C -b localhost:19092,localhost:29092,localhost:39092 -t DemoTopic -p 2
```

#### Publish message to "DemoTopic" partition 2
```sh
echo 'Test call api' | kafkacat -P -b localhost:19092,localhost:29092,localhost:39092 -t DemoTopic -p 2
```

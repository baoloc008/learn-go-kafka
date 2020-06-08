package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"learn-go-kafka/handler"
	"learn-go-kafka/kafka"
	"log"
	"net/http"
	"strings"
)

var (
	listenAddr = ":9620"
	// kafka
	kafkaBrokerUrl = "localhost:19092,localhost:29092,localhost:39092"
	kafkaClientId  = "demo-kafka-client"
	kafkaTopic     = "DemoTopic"
)

func main() {

	// connect to kafka
	kafkaProducer, err := kafka.InitWriter(strings.Split(kafkaBrokerUrl, ","), kafkaClientId, kafkaTopic)
	if err != nil {
		log.Panic("Unable to init kafka writer. ", "Error: ", err.Error())
		return
	}
	defer kafkaProducer.Close()
	kafkaConsumer, err := kafka.InitReader(strings.Split(kafkaBrokerUrl, ","), kafkaClientId, kafkaTopic)
	if err != nil {
		log.Panic("Unable to init kafka reader. ", "Error: ", err.Error())
		return
	}
	defer kafkaConsumer.Close()

	r := echo.New()
	r.Use(middleware.Logger())
	r.GET("/ping", handler.Pong)
	r.POST("/push-message", handler.PushMessage)

	s := http.Server{
		Addr:    listenAddr,
		Handler: r,
	}

	log.Fatal(s.ListenAndServe())
}

package main

import (
	"flag"
	"fmt"
	"testqueue/consumer"
	"testqueue/producer"
)

func main() {
	var isProducer bool
	flag.BoolVar(&isProducer, "p", false, "producer")
	flag.Parse()
	if isProducer {
		// 生产者
		fmt.Println("producer.run")
		producer.Run()
	} else {
		fmt.Println("consumer.run")
		consumer.Run()
	}
}

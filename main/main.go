/**
 * @Time : 2019-07-16 15:16
 * @Author : zhuangjingpeng
 * @File : main
 * @Desc : file function description
 */
package main

import (
	pdmq "GO_PDMQ"
	"fmt"
	"go-nsq"
	"time"
)

type PDMQHandler struct {
}

type NsqHandler struct {
}

func (this *PDMQHandler) HandleMessage(message *pdmq.Message) error {
	fmt.Printf("[PDMQ CONSUMER] [%+v] get handler msg is [%+v]\n", time.Now().Format("2006-01-02 15:04:05"), string(message.Body))
	return nil
}
func (this *NsqHandler) HandleMessage(message *nsq.Message) error {
	fmt.Println(string(message.Body))
	return nil
}
func main() {
	config := pdmq.NewConfig()
	consumer, err := pdmq.NewConsumer("name", "hello", config)
	if err != nil {
		fmt.Println(err)
	}
	consumer.AddHandler(&PDMQHandler{})

	err = consumer.ConnectToPDMQD("127.0.0.1:9400")
	if err != nil {
		fmt.Println(err)
	}

	select {}
}

func main2() {
	consumer, err := nsq.NewConsumer("name", "world", nsq.NewConfig())
	if err != nil {
		fmt.Println(err)
	}
	consumer.AddHandler(&NsqHandler{})
	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		fmt.Println(err)
	}

	select {}
}

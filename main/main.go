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
)

type PDMQHandler struct {
}

func (this *PDMQHandler) HandleMessage(message *pdmq.Message) error {
	fmt.Println(message.Body)
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

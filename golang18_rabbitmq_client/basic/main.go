package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"go-learn-code/golang18_rabbitmq_client/basic/conf"
	"go-learn-code/golang18_rabbitmq_client/basic/entity"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"time"
)
var conn *amqp.Connection
var messageChan chan entity.RabbitmqMessage
var errorSign chan error
var msgs <- chan amqp.Delivery

func main() {
	config:=new(conf.AppConfig)
	yamlFile,err:=ioutil.ReadFile(os.Args[1]+"/config.yaml")

	if err!=nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	yaml.Unmarshal(yamlFile,config)
	messageChan=make(chan entity.RabbitmqMessage,50)
	connectRabbitmqServer(config.RabbitmqAddress,config.RabbitmqQueue)
	for {
		select {
		case err = <- errorSign:
			log.Printf("closing: %s", err)
			log.Printf("reconnect rabbitmq")
			connectRabbitmqServer(config.RabbitmqAddress,config.RabbitmqQueue)
		}
	}
}

func connectRabbitmqServer(address string,queueName string){
	conn,err:=amqp.Dial(address)
	if err!=nil {
		failOnError(err, "fail to connnect to rabbitmq")
	}
	go func() {
		log.Printf("closing: %s", <- conn.NotifyClose(make(chan *amqp.Error)))
		errorSign <- errors.New("Channel Closed")
	}()
	ch, err := conn.Channel()
	err = ch.Qos(1,     // prefetch count //预取数量
		0,     // prefetch size
		false, // global
	)
	// 新建队列，如果这个队列没名字，随机生成一个名字
	q, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	// 队列和交换机绑定，即是队列订阅了发到这个交换机的消息
	//err = ch.QueueBind(q.Name,"","amq.direct",true,nil)
	//failOnError(err, "Failed to bind a queue")
	msgs,err = ch.Consume(q.Name, "", true, false, false, false, nil)
    go handlerMessage(messageChan)
	go handleMQ()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}
func handleMQ(){
	for {
		select {
			case message := <-msgs:
				//message.Ack(false)
				if message.Body != nil && len(message.Body) != 0 {
					log.Printf("received a message : %s", string(message.Body))
					metricsMessage := entity.RabbitmqMessage{};
					json.Unmarshal(message.Body, &metricsMessage)
					//go metricsInfluxDb(metricsMessage)
					if "" != metricsMessage.ApiName {
						messageChan <- metricsMessage
					}
				}
			default:
				//如果上面case都没有成功，则进入default处理流程
				time.Sleep(10 * time.Millisecond)
		}
	}
}
func handlerMessage(messageList chan entity.RabbitmqMessage){
	for message:=range messageList {
		fmt.Println(message)
	}
}
package rabbitmq
//
//import "github.com/streadway/amqp"
//import "log"
//
//
//func GetName() {
//	// 建立连接
//	conn, err := amqp.Dial("amqp://gongyao:gongdandan1209@59.110.225.116:5672/")
//	failOnError(err, "Failed to connect to RabbitMQ")
//	defer conn.Close()
//
//	// 获取channel
//	ch, err := conn.Channel()
//	failOnError(err, "Failed to open a channel")
//	defer ch.Close()
//
//	// 声明队列
//	q, err := ch.QueueDeclare(
//		"kd_sms_send_q", // name
//		true,   // durable
//		false,   // delete when unused
//		false,   // exclusive
//		false,   // no-wait
//		nil,     // arguments
//	)
//	failOnError(err, "Failed to declare a queue")
//
//	// 获取接收消息的Delivery通道
//	msgs, err := ch.Consume(
//		q.Name, // queue
//		"",     // consumer
//		true,   // auto-ack
//		false,  // exclusive
//		false,  // no-local
//		false,  // no-wait
//		nil,    // args
//	)
//	failOnError(err, "Failed to register a consumer")
//
//	forever := make(chan bool)
//
//	go func() {
//		for d := range msgs {
//			log.Printf("Received a message: %s", d.Body)
//		}
//	}()
//
//	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
//	<-forever // 一直阻塞
//
//
//}
//
//func failOnError(err error, msg string) {
//	if err != nil {
//		log.Fatalf("%s: %s", msg, err)
//	}
//}
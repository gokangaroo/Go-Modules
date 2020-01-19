package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

// https://www.rabbitmq.com/getstarted.html
// https://www.cnblogs.com/-wenli/p/12203202.html

// RabbitMQ 声明队列类型
type RabbitMQ struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	Name     string
	exchange string
}

// ReConnect 重连
func (q *RabbitMQ) ReConnect(s string) *RabbitMQ {
	if q.conn != nil {
		e := q.conn.Close()
		failOnError(e, "连接关闭失败！")
	}
	if q.channel != nil {
		e := q.channel.Close()
		failOnError(e, "频道关闭失败！")
	}
	//连接rabbitmq
	conn, e := amqp.Dial(s)
	failOnError(e, "连接Rabbitmq服务器失败！")
	ch, e := conn.Channel()
	failOnError(e, "无法打开频道！")
	mq := new(RabbitMQ)
	mq.conn = conn
	mq.channel = ch
	return mq
}

// New 初始化单个消息队列
//第一个参数：rabbitmq服务器的链接，第二个参数：队列名字
func New(s string, name string) *RabbitMQ {
	//连接rabbitmq
	conn, e := amqp.Dial(s)
	failOnError(e, "连接Rabbitmq服务器失败！")
	ch, e := conn.Channel()
	failOnError(e, "无法打开频道！")
	q, e := ch.QueueDeclare(
		name,  //队列名
		false, //是否开启持久化:保证及时rabbitmq挂了也不会影响=>如果是已有队列, 可能需要删掉再建立
		true,  //不使用时删除
		false, //排他
		false, //不等待
		nil,   //参数
	)
	failOnError(e, "初始化队列失败！")

	mq := new(RabbitMQ)
	mq.conn = conn
	mq.channel = ch
	mq.Name = q.Name
	return mq
}

//批量初始化消息队列
//第一个参数：rabbitmq服务器的链接，第二个参数：队列名字列表

// Qos 配置队列参数: https://www.rabbitmq.com/tutorials/tutorial-two-go.html
func (q *RabbitMQ) Qos() {
	e := q.channel.Qos(1, 0, false)
	failOnError(e, "无法设置QoS")
}

//配置交换机参数

//初始化交换机
//第一个参数：rabbitmq服务器的链接，第二个参数：交换机名字，第三个参数：交换机类型
func NewExchange(s string, name string, typename string) {
	//连接rabbitmq
	conn, e := amqp.Dial(s)
	failOnError(e, "连接Rabbitmq服务器失败！")
	ch, e := conn.Channel()
	failOnError(e, "无法打开频道！")
	e = ch.ExchangeDeclare(
		name,     // name
		typename, // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(e, "初始化交换机失败！")

}

//删除交换机
func (q *RabbitMQ) ExchangeDelete(exchange string) {
	e := q.channel.ExchangeDelete(exchange, false, true)
	failOnError(e, "绑定队列失败！")
}

//绑定消息队列到哪个exchange
func (q *RabbitMQ) Bind(exchange string, key string) {
	e := q.channel.QueueBind(
		q.Name,
		key,
		exchange,
		false,
		nil,
	)
	failOnError(e, "绑定队列失败！")
	q.exchange = exchange
}

//向消息队列发送消息
//Send方法可以往某个消息队列发送消息
func (q *RabbitMQ) Send(body interface{}) {
	str, e := json.Marshal(body)
	failOnError(e, "消息序列化失败！")
	e = q.channel.Publish(
		"",     //交换
		q.Name, //路由键：当前队列的名字
		false,  //必填
		false,  //立即
		amqp.Publishing{
			ReplyTo: q.Name,
			Body:    []byte(str),
		})
	msg := "向队列:" + q.Name + "发送消息失败！"
	failOnError(e, msg)
}

//向exchange发送消息
//Publish方法可以往某个exchange发送消息
func (q *RabbitMQ) Publish(exchange string, body interface{}, key string) {
	str, e := json.Marshal(body)
	failOnError(e, "消息序列化失败！")
	e = q.channel.Publish(
		exchange,
		key,
		false,
		false,
		amqp.Publishing{ReplyTo: q.Name,
			Body: []byte(str)},
	)
	failOnError(e, "向路由发送消息失败！")
}

//接收某个消息队列的消息
func (q *RabbitMQ) Consume() <-chan amqp.Delivery {
	c, e := q.channel.Consume(
		q.Name, //指定从哪个队列中接收消息
		"",
		false, // 设置为false表示需要我们手动ack
		false,
		false,
		false,
		nil,
	)
	failOnError(e, "接收消息失败！")
	return c
}

//关闭队列连接
func (q *RabbitMQ) Close() {
	q.channel.Close()
}

//错误处理函数
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

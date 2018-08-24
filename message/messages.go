package nats


//nats 사용되는 메시지 인터페이스

type Tag interface {
	getName() string
}


type message interface {
	MsgTopic() string
	MsgData() string
	Tags() []string
}

package nats


//nats 사용되는 메시지 인터페이스

type message interface {
	MsgTopic() string
	MsgContent() string
}

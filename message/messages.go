package nats


//nats 사용되는 메시지 인터페이스

type Tag interface {
	getName() string
	setName(AName string)
	getVaue() {}
	setValue(Value {})
}


type message interface {
	getDeviceId() string
	setDeviceId(deviceId string)
	Tags() []Tag
}

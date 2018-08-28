package nats

import (
	"time"
)


//nats 사용되는 메시지 인터페이스
type ValueType interface {
	getValue()
	
}

type Tag struct {
	TagName string
	Value ValueType
}


type message struct {
	DeviceId string
	Tags []*Tag
	RevTime time
}

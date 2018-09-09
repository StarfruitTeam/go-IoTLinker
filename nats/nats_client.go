package nats

import (
	"github.com/nats-io/go-nats"
	"fmt"
	"log"
)
/*
목적
1. nats 의 메시지를 수신한다. 
2. nats 에 메시지즐 발행한다. 
*/
// Nats 용 클라이언트


type nats_client struct {
	URL string // nats 의 주소
	NatsClient nats.Conn
}


func (nats_client) Connect(token,urls string):boolean{
	//nats 서버와 연결합니다.
	//연결에 성공하면 Connected 를 변경합니다.
	nc,err := nats.Connect(*urls)
	if err != nil {
		log.Fatalf("Can't connect: %v\n",err)
	}
	NatsClient = nc; 
	
}

func (nats_client) Publish(ClientId,topic,msg string):boolean{
	//nats 서버에 메시지를 발행합니다.
	//체널의 메시지를 발행합니다.
	//
	NatsClient.Publish(topic,msg []byte)
	NatsClient.Flush()

	if err := NatsClient.LastError(); err != nil {
		log.Fatal(err)
	} else {

	}
}


func (nats_clinet) Subscribe(ClientId string,topic string):boolean{
	//nats 서버의 메시지를 수신한다.
	//수신한 메시지를 체널에 넣습니다.
	//개별 토픽에 대한 구독 요청
	//MQTT에 대응하는 규칙에 대한 함수 추가 필요함
	NatsClient.Subscribe(topic func(msg *nats.Msg) {
		// Subscribe 추가 
	})
	NatsClient.Flush()

	if err := NatsClient.LastError(); err != nil {
		log.Fatal(err)
	}

	
}

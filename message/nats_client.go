package nats

import (
	"github.com/nats-io/go-nats"
)
/*
목적
1. nats 의 메시지를 수신한다. 
2. nats 에 메시지즐 발행한다. 
*/
// Nats 용 클라이언트 
type nats_client struct {
	URL string // nats 의 주소
	Connected boolean // 연결 여부 
	done chan int   // 작업 종료 여부
	rev_queue chan string // Recevie 메시지 큐
	send_queue chan string // Send 메시지 큐
}


func (nats_client) Connect(token string):boolean{
	//nats 서버와 연결합니다.
	//연결에 성공하면 Connected 를 변경합니다. 

}

func (nats_client) Publish(ClientId string,msg string):boolean{
	//nats 서버에 메시지를 발행합니다.
	//체널의 메시지를 발행합니다. 
}


func (nats_clinet) Subscribe(ClientId string,tag string):boolean{
	//nats 서버의 메시지를 수신한다.
	//수신한 메시지를 체널에 넣습니다. 
}

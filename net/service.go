dpackage NetService

import "../nats/nats_client"
import "./receiver"

type NetServcie struct {
	receivers map[string]receiver
	buffers chan
	client nats_client
}

// receiver 에서 수집한 메시지는 nats_client 에서 nats Stream로 전송
func (*NetService)AddReceiver(name string;r receiver){
	receivers[name] = receiver
}

// 서비스 생성
func (*NatService) NewService(*c nats_client){
	receivers := make(map[string]receiver)
	buffers = meke(chan message)
	client = c;
}


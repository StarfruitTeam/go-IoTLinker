package NetService

//MQTT, Socket, RestApi,CoAP 용 커넥터 인터페이스

type receiver interfce {
	runRecevicer(name string;port int;buffer chan)
}


type mqttReceiver struct {
	name string
	buffer chan
	port int
	client mqtt_client
}

func (r *mqttReceiver) runRecevicer(name string;port int;buffer chan){
	// mqtt client 생성
	// go_routine 생성
	// 각 프로토콜에 따라서 메시지즐 받아서 buffer 로 전달	
}






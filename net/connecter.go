package NetService

//MQTT, Socket, RestApi,CoAP 용 커넥터 인터페이스

type connecter interfce {
	SetConfig(name string,port int)
	SetNatClient(client nat_client)
}

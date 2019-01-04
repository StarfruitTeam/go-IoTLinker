package net

/*
기능 정의 
1. 프로토콜에서 데이터를 수신한다. 수신 방법은 프로토콜의 사양을 따른다. 
2. 받은 메시지를 messageing 시스템에 발행한다. 
3. 메시징 시스템의 메시지를 수신하여 클라이언트로 전송(이벤트 메시지 , 행위 메시지)
*/
type MessageConnector interface {
	Connect(id,url string) boolean	// 메시징 시스템에 연결 한다. 
	PubMsg(msg string) boolean // 발행 기능 
	SubMsg(*topic,byte(qos),queue chan<- {}) //구독 기능 조건: 토픽을 지정한다.	
}
/*
    """
    recevie mqtt messag
    """
    message = dict(
        topic=message.topic,
        data=message.payload.decode(),
        datetime=datetime.now()
    )
    """
    topic 의 정보를 이용해서 tag 정보를 정한다.
    """
    mqtt_process.process(message)

*/

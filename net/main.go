package message_gateway

import (
	"fmt"
	"flag"
	"os"
	"context"
	"../message/message"
	""
)

/*
 Message Gateway
1. MQTT Pub/Sub 기능 추가 
2. nats 연결 기능 추가 

*/
func main() {
	topic := flag.String("topic", "", "The topic name to/from which to publish/subscribe")
	broker := flag.String("broker", "tcp://127.0.0.1:1883", "The broker URI. ex: tcp://10.10.1.1:1883")
	password := flag.String("password", "", "The password (optional)")
	user := flag.String("user", "", "The User (optional)")
	id := flag.String("id", "testgoid", "The ClientID (optional)")
	cleansess := flag.Bool("clean", false, "Set Clean Session (default false)")
	qos := flag.Int("qos", 0, "The Quality of Service 0,1,2 (default 0)")
	num := flag.Int("num", 1, "The number of messages to publish or subscribe (default 1)")
	payload := flag.String("message", "", "The message text to publish (default empty)")
	action := flag.String("action", "", "Action publish or subscribe (required)")
	store := flag.String("store", ":memory:", "The Store Directory (default use memory store)")
	nats := flag.String("nats", ":127.0.0.1:4222:", "The NATS Messaging System (default:127.0.0.1:4222)")

	//mongoDB 연결 정보를 

	flag.Parse()
	//1. 메시지를 받는다 
	//2. topic 정보를 조회 하여 디바이스를 찾는다.
	//3. 공용 메시지 구조로 변환한다.
	/*
공용 메시지 구조
- DEVICE_ID:UUID
- TAG_DATAS
  - TAG:DATA
  - TAG:DATA
   
*/

	
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

class MqttMessageProcess(object):
    contextEngin = None
    
    def __init__(self):
        self.contextEngin = ContextEngin()

    def finderTag(self, device_uuid, tag_name):
        print(device_uuid)
        result = None
        result_device = Device.objects(uuid=device_uuid,
                                       tags__name=tag_name).first()
        for tag in result_device.tags:
            if tag.name == tag_name:
                result = tag
        return result

    def process(self, message):
        """
        1. 룰이 걸려 있는 테크를 찾는다.
        2. 데이터 처리 프로세스를 수행한다.
        3. custom topic 처리 redis 저장된 토픽 처리
        """
        device = loads(redis_store.get(message['topic']))
        tag = self.finderTag(device['device_uuid'],
                             device['tag_name'])
        if self.contextEngin is not None:
            self.contextEngin.start(tag, message)
        else:
            print('contextEngin is None')

        self.saveTagData(tag, device['device_uuid'], message)

    def saveTagData(self, tag, device_uuid, message):
        tagData = TagData(tag=tag.name, device_uuid=device_uuid,
                          value=message['data'], data_type=tag.data_type)
        tagData.save()
        """
        tag 정보를 데이터 베이스에 저장
        """
       

*/




	
}

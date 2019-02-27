package NetService
import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

//MQTT, Socket, RestApi,CoAP 용 커넥터 인터페이스

type config struct {
	broker string //bloker
	password string //"tcp://127.0.0.1:1883", "The broker URI. ex: tcp://10.10.1.1:1883"
	user string //The password (optional)
	id string //The ClientID (optional) 
	cleansess boolean //Set Clean Session (default false)
	qos int //The Quality of Service 0,1,2 (default 0)
	num int //"The number of messages to publish or subscribe (default 1)
	payload string//The message text to publish (default empty)
	action string //Action publish or subscribe (required)
	store string//The Store Directory (default use memory store)
}

type receiver interfce {
	runRecevicer(cfg *config;buffer chan)
	Subscribe(AData string[])
}


type mqttReceiver struct {
	name string
	buffer chan
	port int
	client mqtt_client
}

func (r *mqttReceiver) runRecevicer(cfg *config;buffer chan){

	
	opts := MQTT.NewClientOptions()
	opts.AddBroker(*cfg.broker)
	opts.SetClientID(*cfg.id)
	opts.SetUsername(*cfg.user)
	opts.SetPassword(*cfg.password)
	opts.SetCleanSession(*cfg.cleansess)

	if *cfg.store != ":memory:" {
		opts.SetStore(MQTT.NewFileStore(*store))
	}

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe(*topic, byte(*qos), nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for receiveCount < *num {
		incoming := <-choke
		fmt.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
		receiveCount++
	}


	client.Disconnect(250)
	fmt.Println("Sample Publisher Disconnected")

	
	go func (){
		for(){
			// 
		}
	}
	// go_routine 생성
	// 서브스크립션 토픽 등록 
	// 
	// 각 프로토콜에 따라서 메시지즐 받아서 buffer 로 전달	
}

func (r *mqttReceiver) subscribe(AData string[]){
	// 슬라이스의 토픽 구독 설정
}






package main

 import (
	  MQTT "github.com/eclipse/paho.mqtt.golang"
	 "flag"
	 "fmt"
	 "time"
	 "os"
 )

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC:%s\n",msg.Topic())
	fmt.Printf("TOPIC:%s\n",msg.MessageID())
}

func main(){
	//topic := flag.String("topic", "", "The topic name to/from which to publish/subscribe")
	broker := flag.String("broker", "tcp://127.0.0.1:1883", "The broker URI. ex: tcp://10.10.1.1:1883")
	//password := flag.String("password", "", "The password (optional)")
	//user := flag.String("user", "", "The User (optional)")
	//id := flag.String("id", "testgoid", "The ClientID (optional)")
	//cleansess := flag.Bool("clean", false, "Set Clean Session (default false)")
	//qos := flag.Int("qos", 0, "The Quality of Service 0,1,2 (default 0)")
	//num := flag.Int("num", 1, "The number of messages to publish or subscribe (default 1)")
	//payload := flag.String("message", "", "The message text to publish (default empty)")
	//action := flag.String("action", "", "Action publish or subscribe (required)")
	//store := flag.String("store", ":memory:", "The Store Directory (default use memory store)")
	flag.Parse()

	opts := MQTT.NewClientOptions().AddBroker(*broker)

	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)


	client := MQTT.NewClient(opts)

	if token := client.Connect();token.Wait() && token.Error() != nil {
	panic(token.Error())
	}

	if token := client.Subscribe("go-mqtt/sample", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := client.Publish("go-mqtt/sample", 0, false, text)
		token.Wait()
	}

	time.Sleep(6 * time.Second)

	if token := client.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	client.Disconnect(250)

	time.Sleep(1 * time.Second)
}



	
}

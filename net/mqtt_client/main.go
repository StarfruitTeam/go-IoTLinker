package main

 import (
	  MQTT "github.com/eclipse/paho.mqtt.golang"
	 "flag"
	 "fmt"
 )

func main(){
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
	flag.Parse()



	flag.Parse()
	options := mqtt.NewClientOptions()
	append(options.Servers,"127.0.0.")
	options.

	var client = nil
	client = mqtt.NewClient()


	fmt.Printf("Hola %s!\n", *topic)

	
}
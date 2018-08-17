package nats

import (
	"github.com/nats-io/go-nats"
)


// Nats 용 클라이언트 
type nats_client struct {
	URL string
	Connected boolean
	Done chan int
	buf chan string	
}


func (nats_client) Publish(ClientId string,msg string):boolean{

}


func (nats_clinet) Subscribe(ClientId string,tag string):boolean{

}
